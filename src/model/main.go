package model

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/tokopedia/wednesday/src/model/booking"
	mock_booking "github.com/tokopedia/wednesday/src/model/booking/mocks"
	"github.com/tokopedia/wednesday/src/model/cab"
	"github.com/tokopedia/wednesday/src/model/user"
)

type Models struct {
	User    user.User
	Cab     cab.Cab
	Booking booking.Booking
}

func Init() Models {
	return Models{
		User:    user.Init(),
		Cab:     cab.Init(),
		Booking: booking.Init(),
	}
}
func InitMockModule() Models {
	t := testing.T{}
	mockCtrl := gomock.NewController(&t)
	defer mockCtrl.Finish()

	bookingRepo := mock_booking.NewMockBooking(mockCtrl)
	bookingRepo.EXPECT().GetLastNBookings(context.Background(), 1, 3).Return([]booking.Info{
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
	})
	return Models{
		Booking: bookingRepo,
	}
}
