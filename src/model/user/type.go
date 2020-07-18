package user

import "context"

type User interface {
	GetLastBookings(ctx context.Context, user UserInfo) ([]User, error)
}

type UserInfo struct {
	ID   int64  `json:"id"`
	Name string `json"name"`
}

type repo struct {
}
