package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tokopedia/wednesday/src/apperror"
	"github.com/tokopedia/wednesday/src/model"
	"github.com/tokopedia/wednesday/src/server"
)

func Init(models model.Models) Handller {
	return repo{
		models: models,
	}
}

func (h repo) GetNearByCabs(ctx context.Context, r *http.Request, p server.HttpParams) (result interface{}, err error) {
	return
}

func (h repo) GetBookings(ctx context.Context, r *http.Request, p server.HttpParams) (result interface{}, err error) {
	var (
		payload LastBookingListPayload = LastBookingListPayload{}
	)

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("Inavlid Payload %+v", err)
		err = apperror.BadError("INV_PAYLOAD", "Request payload is Invalid")
		return
	}
	result, err = h.models.Booking.GetLastNBookings(ctx, payload.UserID, payload.Limit)
	return
}
