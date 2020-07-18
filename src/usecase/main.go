package usecase

import (
	"github.com/tokopedia/wednesday/src/usecase/user"
)

type Usecases struct {
	User user.User
}

func Init() Usecases {
	return Usecases{
		User: user.Init(),
	}
}
