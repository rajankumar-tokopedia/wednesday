package handler

import (
	"context"
	"net/http"

	"github.com/tokopedia/wednesday/src/server"
	"github.com/tokopedia/wednesday/src/usecase"
)

type Handller interface {
	GetNearByCabs(ctx context.Context, r *http.Request, p server.HttpParams) (interface{}, error)
	GetBookings(ctx context.Context, r *http.Request, p server.HttpParams) (interface{}, error)
	BookCab(ctx context.Context, r *http.Request, p server.HttpParams) (result interface{}, err error)
}

type repo struct {
	usecases usecase.Usecases
}

type LastBookingListPayload struct {
	UserID int64 `json:"user_id"`
	Limit  int64 `json:"limit"`
}

type BookCabPayload struct {
	From   string `json:"from"`
	To     string `json:"to"`
	UserID int64  `json:"user_id"`
}
type NearByCabsPayload struct {
	Lon string `json:"lon"`
	Lat string `json:"lat"`
}
