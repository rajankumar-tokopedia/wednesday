package server

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	apperr "github.com/tokopedia/wednesday/src/error"
)

type repo struct {
	error       apperr.AppError
	router      *httprouter.Router
	middlewares []HttpMiddleware
}

type HttpParams interface {
	ByName(string) string
}

type ResponseHeader struct {
	SPT       string `json:"process_time"`
	ST        string `json:"server_time"`
	ErrCode   string `json:"error_code,omitempty"`
	StausCode int    `json:"status_code"`
}

type Response struct {
	Header    ResponseHeader `json:"header"`
	Data      interface{}    `json:"data"`
	startTime time.Time
}

type HttpHandler func(context.Context, *http.Request, HttpParams) (interface{}, error)

type HttpMiddleware func(context.Context, []HttpMiddleware, *http.Request, HttpParams) (interface{}, error)

type HttpServer interface {
	GET(string, HttpHandler)
	POST(string, HttpHandler)
	PUT(string, HttpHandler)
	DELETE(string, HttpHandler)
	OPTIONS(string, HttpHandler)

	ServeHTTP(http.ResponseWriter, *http.Request)
	ServeFiles(string, http.FileSystem)

	AddMiddleware(HttpMiddleware)
}
