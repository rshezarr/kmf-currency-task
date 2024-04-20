package model

type CustomError struct {
	StatusCode  int
	RespMessage string
	Error       error
}

func NewError(statCode int, respMsg string, err error) *CustomError {
	return &CustomError{
		StatusCode:  statCode,
		RespMessage: respMsg,
		Error:       err,
	}
}
