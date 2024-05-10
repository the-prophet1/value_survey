package errors

type Error interface {
	error
	Code() int
	Message() string
	StackTrace() string
}
