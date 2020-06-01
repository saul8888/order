package errors

import "github.com/labstack/echo"

type Error struct {
	IsTrue  bool        `json:"error"`
	Message interface{} `json:"message"`
}

func New(message string) Error {
	return Error{
		IsTrue:  true,
		Message: message,
	}
}

func NewError(err error) Error {
	er := Error{}

	er.IsTrue = true

	switch v := err.(type) {
	case *echo.HTTPError:
		er.Message = v.Message
	default:
		er.Message = v.Error()
	}

	return er
}
