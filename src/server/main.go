package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/tokopedia/wednesday/src/apperror"
)

func next(ctx context.Context, httpHandler HttpHandler, server *repo, r *http.Request, p HttpParams) (interface{}, error) {
	lastMiddleware := func(ctx context.Context, middlewares []HttpMiddleware, r *http.Request, p HttpParams) (interface{}, error) {
		return httpHandler(ctx, r, p)
	}
	middlewares := server.middlewares
	middlewares = append(middlewares, lastMiddleware)

	return middlewares[0](ctx, middlewares[1:], r, p)
}

func (s *repo) routerHandler(httpHandler HttpHandler, server *repo) httprouter.Handle {
	routeHandler := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		response := Response{
			Header: ResponseHeader{
				StausCode: 200,
			},
		}

		response.StartProcessing()
		responseBody, err := next(context.Background(), httpHandler, server, r, p)

		w.Header().Set("Content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if err != nil {
			errcode, msg, status := s.error.ErrorHandler(err)
			response.Header.ErrCode = errcode
			response.Header.StausCode = status
			response.Header.Message = msg
			w.WriteHeader(status)
			response.EndProcessing()
			errorBody, _ := json.Marshal(response)
			w.Write(errorBody)
			return
		}

		response.EndProcessing()
		response.Data = responseBody
		resByte, err := json.Marshal(response)
		if err != nil {
			errcode, msg, status := s.error.ErrorHandler(err)
			response.Header.ErrCode = errcode
			response.Header.StausCode = status
			response.Header.Message = msg
			w.WriteHeader(status)
			response.EndProcessing()
			return
		}

		_, err = w.Write(resByte)
		if err != nil {
			log.Printf("Failed to write response %+v", err)
		}
	}

	return routeHandler
}

func NewHttpServer(appErrorRepo apperror.AppError) HttpServer {
	router := httprouter.New()
	return &repo{
		router: router,
		error:  appErrorRepo,
	}
}

func (s *repo) GET(path string, httpHandler HttpHandler) {
	routeHandler := s.routerHandler(httpHandler, s)
	s.router.GET(path, routeHandler)
}

func (s *repo) POST(path string, httpHandler HttpHandler) {
	routeHandler := s.routerHandler(httpHandler, s)
	s.router.POST(path, routeHandler)
}

func (s *repo) PUT(path string, httpHandler HttpHandler) {
	routeHandler := s.routerHandler(httpHandler, s)
	s.router.PUT(path, routeHandler)
}

func (s *repo) DELETE(path string, httpHandler HttpHandler) {
	routeHandler := s.routerHandler(httpHandler, s)
	s.router.DELETE(path, routeHandler)
}

func (s *repo) OPTIONS(path string, httpHandler HttpHandler) {
	routeHandler := s.routerHandler(httpHandler, s)
	s.router.OPTIONS(path, routeHandler)
}
func (s *repo) AddMiddleware(middleware HttpMiddleware) {
	s.middlewares = append(s.middlewares, middleware)
}

func (s *repo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *repo) ServeFiles(path string, root http.FileSystem) {
	s.router.ServeFiles(path, root)
}

func (resp *Response) StartProcessing() {
	resp.startTime = time.Now()
	return
}
func (resp *Response) EndProcessing() {
	processTime := time.Since(resp.startTime).Seconds() * 1000
	resp.Header.SPT = strconv.FormatFloat(processTime, 'f', -1, 64)
	resp.Header.ST = strconv.FormatInt(resp.startTime.Unix(), 10)
	return
}
