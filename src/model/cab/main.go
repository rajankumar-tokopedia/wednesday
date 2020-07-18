package cab

import "context"

func Init() Cab {
	return repo{}
}

func (r repo) NearBy(ctx context.Context, lon, lat string) (i []Info, err error) {
	return
}
