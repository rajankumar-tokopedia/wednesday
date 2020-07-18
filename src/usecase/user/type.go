package user

import (
	"context"

	"github.com/tokopedia/wednesday/src/model"
)

type User interface {
	InfoById(ctx context.Context, userId int64) ([]model.UserInfo, error)
	GetLastNBookings(ctx context.Context, userId int64, n int64) (i []model.BookingInfo, err error)
	NearByCabs(ctx context.Context, lon, lat string) (i []model.CabInfo, err error)
	BookCab(ctx context.Context, from, to string, userId int64) (i model.BookingInfo, err error)
}

type repo struct {
}
