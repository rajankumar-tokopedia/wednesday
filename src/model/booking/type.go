package booking

import (
	"context"
	"time"
)

type Booking interface {
	GetLastNBookings(ctx context.Context, userId int64, n int64) (i []Info, err error)
}

type Info struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
	CabID  int64 `json:"cab_id"`

	Time     time.Time
	From     string `json:"from"`
	To       string `json:"to"`
	Fare     int64  `json:"fare"`
	Currency string `json:"currency"`
}

type repo struct {
}
