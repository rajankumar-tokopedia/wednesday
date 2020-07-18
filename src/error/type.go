package error

type AppError interface {
	ErrorHandler(error) ([]byte, int)
}

type repo struct {
}
