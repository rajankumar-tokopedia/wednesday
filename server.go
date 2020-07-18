package main

import (
	"context"
	"log"
	"net/http"

	"gopkg.in/paytm/grace.v1"

	"github.com/tokopedia/wednesday/src/constants"
	errorRepo "github.com/tokopedia/wednesday/src/error"
	"github.com/tokopedia/wednesday/src/middleware"
	"github.com/tokopedia/wednesday/src/server"
)

func main() {
	// server and http routes
	appError := errorRepo.Init()
	appServer := server.NewHttpServer(appError)

	appServer.AddMiddleware(middleware.SetupContext)
	//Healthcheck handler
	appServer.GET("/healthcheck", func(ctx context.Context, r *http.Request, p server.HttpParams) (interface{}, error) {
		return map[string]string{
			"ping": "pong",
		}, nil
	})

	log.Fatal(grace.Serve(constants.Port, appServer))
}
