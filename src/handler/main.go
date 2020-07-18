package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tokopedia/wednesday/src/apperror"
	"github.com/tokopedia/wednesday/src/server"
	"github.com/tokopedia/wednesday/src/usecase"
)

func Init(usecase usecase.Usecases) Handller {
	return repo{
		usecases: usecase,
	}
}

func (h repo) GetNearByCabs(ctx context.Context, r *http.Request, p server.HttpParams) (result interface{}, err error) {

	var (
		payload NearByCabsPayload = NearByCabsPayload{}
	)

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("Inavlid Payload %+v", err)
		err = apperror.BadError("INV_PAYLOAD", "Request payload is Invalid")
		return
	}
	result, err = h.usecases.User.NearByCabs(ctx, payload.Lon, payload.Lat)
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
	result, err = h.usecases.User.GetLastNBookings(ctx, payload.UserID, payload.Limit)
	return
}

func (h repo) BookCab(ctx context.Context, r *http.Request, p server.HttpParams) (result interface{}, err error) {
	var (
		payload BookCabPayload = BookCabPayload{}
	)

	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("Inavlid Payload %+v", err)
		err = apperror.BadError("INV_PAYLOAD", "Request payload is Invalid")
		return
	}
	result, err = h.usecases.User.BookCab(ctx, payload.From, payload.To, payload.UserID)
	return
}
