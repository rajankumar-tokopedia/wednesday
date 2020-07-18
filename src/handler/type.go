package handler

import (
	"context"
	"net/http"

	"github.com/tokopedia/wednesday/src/model"
	"github.com/tokopedia/wednesday/src/server"
)

type Handller interface {
	GetNearByCabs(ctx context.Context, r *http.Request, p server.HttpParams) (interface{}, error)
	GetBookings(ctx context.Context, r *http.Request, p server.HttpParams) (interface{}, error)
}

type repo struct {
	models model.Models
}

type LastBookingListPayload struct {
	UserID int64 `json:"user_id"`
	Limit  int64 `json:"limit"`
}
