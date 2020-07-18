package user

import "context"

func Init() User {
	return repo{}
}

func (r repo) GetLastBookings(ctx context.Context, user UserInfo) (users []User, err error) {
	return
}
