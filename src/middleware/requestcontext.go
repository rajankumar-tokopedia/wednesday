package middleware

import (
	"context"
	"net/http"

	"github.com/tokopedia/wednesday/src/constants"
	"github.com/tokopedia/wednesday/src/server"
)

func SetupContext(ctx context.Context, next []server.HttpMiddleware, r *http.Request, p server.HttpParams) (interface{}, error) {

	ctx = context.WithValue(ctx, "App-Name", constants.AppName)
	return next[0](ctx, next[1:], r, p)
}
