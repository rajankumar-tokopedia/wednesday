package booking

import (
	"context"
	"time"
)

func Init() Booking {
	return repo{}
}

func (r repo) GetLastNBookings(ctx context.Context, userId int64, n int64) (i []Info, err error) {
	i = []Info{
		{
			ID:       1,
			UserID:   1,
			CabID:    1,
			Time:     time.Now().Add(-12 * time.Hour),
			From:     "A",
			To:       "B",
			Fare:     10,
			Currency: "INR",
		},
		{
			ID:       2,
			UserID:   1,
			CabID:    2,
			Time:     time.Now().Add(-15 * time.Hour),
			From:     "K",
			To:       "A",
			Fare:     30,
			Currency: "INR",
		},
		{
			ID:       3,
			UserID:   1,
			CabID:    1,
			Time:     time.Now().Add(-16 * time.Hour),
			From:     "K",
			To:       "J",
			Fare:     50,
			Currency: "INR",
		},
	}
	return
}
