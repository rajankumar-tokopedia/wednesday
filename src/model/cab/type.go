package cab

import "context"

type Cab interface {
	NearBy(ctx context.Context, lon, lat string) (i []Info, err error)
}

type Info struct {
	ID       int64   `json:"id"`
	Number   string  `json:"number"`
	Distance float64 `json:"distance"`
}

type repo struct {
}
