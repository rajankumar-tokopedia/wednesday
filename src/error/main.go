package error

import (
	"encoding/json"
	"fmt"
	"log"
)

type Error struct {
	code     string
	message  string
	httpCode int
}

func BadError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 400,
	}
}

func NotFoundError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 404,
	}
}

func InternalServerError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 500,
	}
}

func ForbiddenError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 403,
	}
}

func NotImplementedError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 501,
	}
}

func BadGatewayError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 502,
	}
}

func HttpErrorf(format string, a ...interface{}) string {
	errmsg := fmt.Sprintf(format, a...)
	return fmt.Sprintf(`{"success": false,"message_error": "%s"}`, errmsg)
}

func (err *Error) Error() string {
	return err.message
}

func (err *Error) Code() string {
	return err.code
}

func (err *Error) HttpCode() int {
	return err.httpCode
}

func (r repo) ErrorHandler(err error) ([]byte, int) {
	var (
		errResp      interface{}
		responseBody []byte
	)
	appErr, ok := err.(*Error)

	if !ok {
		appErr = InternalServerError(err.Error(), "Something went wrong, please try after sometime")
	}

	errInUnMarshal := json.Unmarshal([]byte(appErr.Error()), &errResp)
	if errInUnMarshal != nil {
		log.Printf("Unable to unmarshal err string %+v", err)
		responseBody, _ = json.Marshal(map[string]interface{}{
			"code":    appErr.Code(),
			"message": appErr.Error(),
		})

	} else {
		responseBody, _ = json.Marshal(map[string]interface{}{
			"code":    appErr.Code(),
			"message": errResp,
			"err":     errResp,
		})
	}
	return responseBody, appErr.HttpCode()
}
func Init() AppError {
	return repo{}
}
