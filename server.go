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
	"github.com/tokopedia/wednesday/src/model"
	"github.com/tokopedia/wednesday/src/server"
)

func main() {
	// server and http routes
	appError := apperror.Init()
	appServer := server.NewHttpServer(appError)

	appServer.AddMiddleware(middleware.SetupContext)
	models := model.Init()
	//models := model.InitMockModule()
	h := handler.Init(models)
	//Healthcheck handler
	appServer.GET("/healthcheck", func(ctx context.Context, r *http.Request, p server.HttpParams) (interface{}, error) {
		return map[string]string{
			"ping": "pong",
		}, nil
	})

	//	curl --location --request POST 'localhost:5176/api/v1/bookings' \
	//	--header 'Content-Type: application/json' \
	//	--data-raw '{
	//	"user_id":1,
	//		"limit":3
	//}'
	appServer.POST("/api/v1/bookings", h.GetBookings)

	log.Fatal(grace.Serve(constants.Port, appServer))
}
