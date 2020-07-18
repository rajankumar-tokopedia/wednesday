package main

import (
	"context"
	"log"
	"net/http"

	"gopkg.in/paytm/grace.v1"

	"github.com/tokopedia/wednesday/src/apperror"
	"github.com/tokopedia/wednesday/src/constants"
	"github.com/tokopedia/wednesday/src/handler"
	"github.com/tokopedia/wednesday/src/middleware"
	"github.com/tokopedia/wednesday/src/server"
	"github.com/tokopedia/wednesday/src/usecase"
)

func main() {
	// server and http routes
	appError := apperror.Init()
	appServer := server.NewHttpServer(appError)

	appServer.AddMiddleware(middleware.SetupContext)
	models := usecase.Init()
	//models := model.InitMockModule()
	h := handler.Init(models)
	//Healthcheck handler
	appServer.GET("/healthcheck", func(ctx context.Context, r *http.Request, p server.HttpParams) (interface{}, error) {
		return map[string]string{
			"ping": "pong",
		}, nil
	})

	appServer.POST("/api/v1/user/bookings", h.GetBookings)
	appServer.POST("/api/v1/cabs/near", h.GetNearByCabs)
	appServer.POST("/api/v1/cab/book", h.BookCab)

	log.Fatal(grace.Serve(constants.Port, appServer))
}
