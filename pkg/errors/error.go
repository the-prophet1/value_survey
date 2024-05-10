package errors

import "runtime/debug"

type defaultError struct {
	code       int
	message    string
	stackTrace string
}

func (d *defaultError) Error() string {
	return d.message
}

func (d *defaultError) Code() int {
	return d.code
}

func (d *defaultError) Message() string {
	return d.message
}

func (d *defaultError) StackTrace() string {
	return d.stackTrace
}

func New(code int, message string) Error {
	return &defaultError{
		code:       code,
		message:    message,
		stackTrace: string(debug.Stack()),
	}
}

func NewWithError(code int, err error) Error {
	if err == nil {
		return nil
	}
	return &defaultError{
		code:       code,
		message:    err.Error(),
		stackTrace: string(debug.Stack()),
	}
}
