package internal_error

import "net/http"

type InternalError struct {
	Message string
	Err     string
	ErrCode int
}

func (ie *InternalError) Error() string {
	return ie.Message
}

func HttpError(statusCode int) *InternalError {
	if statusCode == 0 {
		return NewInternalServerError("something went wrong")
	}
	switch statusCode {
	case 422:
		return NewBadRequestError(" invalid zipcode")
	case 404:
		return NewNotFoundError("can not find zipcode")
	default:
		return NewInternalServerError("something went wrong")
	}
}
func NewBadRequestError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "invalid zipcode",
		ErrCode: http.StatusUnprocessableEntity,
	}
}

func NewNotFoundError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "not_found",
		ErrCode: http.StatusNotFound,
	}
}

func NewInternalServerError(message string) *InternalError {
	return &InternalError{
		Message: message,
		Err:     "internal_server_error",
		ErrCode: http.StatusInternalServerError,
	}
}
