package model

import "time"

type UserInfo struct {
	ID   int64  `json:"id"`
	Name string `json"name"`
}

type CabInfo struct {
	ID       int64   `json:"id"`
	Number   string  `json:"number"`
	Distance float64 `json:"distance"`
	Unit     string  `json:"unit"`
}

type BookingInfo struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
	CabID  int64 `json:"cab_id"`

	Time     time.Time
	From     string `json:"from"`
	To       string `json:"to"`
	Fare     int64  `json:"fare"`
	Currency string `json:"currency"`
	State    string `json:"state"`
}
