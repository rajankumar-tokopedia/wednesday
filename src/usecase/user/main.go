package user

import (
	"context"
	"time"

	"github.com/tokopedia/wednesday/src/constants"
	"github.com/tokopedia/wednesday/src/model"
)

func Init() User {
	return repo{}
}

func (r repo) InfoById(ctx context.Context, user int64) (users []model.UserInfo, err error) {
	return
}

func (r repo) GetLastNBookings(ctx context.Context, userId int64, n int64) (i []model.BookingInfo, err error) {
	//TODO :: Insted returning this data should implement db layer in  model package and get data
	i = []model.BookingInfo{
		{
			ID:       1,
			UserID:   1,
			CabID:    1,
			Time:     time.Now().Add(-12 * time.Hour),
			From:     "A",
			To:       "B",
			Fare:     10,
			Currency: "INR",
			State:    constants.BookingState.Completed,
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
			State:    constants.BookingState.Completed,
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
			State:    constants.BookingState.Failed,
		},
	}
	return
}

func (r repo) NearByCabs(ctx context.Context, lon, lat string) (i []model.CabInfo, err error) {
	//TODO :: Insted returning this data should implement db layer in  model package and get data
	i = []model.CabInfo{
		{
			1,
			"UP-14-DS-6246",
			100,
			"m",
		},
		{
			2,
			"UP-14-DS-6245",
			0.5,
			"km",
		},
		{
			3,
			"UP-14-DS-6243",
			400,
			"m",
		},
	}
	return
}

func (r repo) BookCab(ctx context.Context, from, to string, userId int64) (i model.BookingInfo, err error) {
	//TODO :: Insted returning this data should implement db layer in  model package and get data
	i = model.BookingInfo{
		ID:       3,
		UserID:   1,
		CabID:    1,
		Time:     time.Now().Add(-16 * time.Hour),
		From:     "K",
		To:       "J",
		Fare:     50,
		Currency: "INR",
		State:    constants.BookingState.Booked,
	}
	return
}
