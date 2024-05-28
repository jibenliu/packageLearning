package main

const (
	// 10000 ~ 19999 validation errors
	INVALID_PARAM int = 10000
	NO_RECORD     int = 10001

	// 20000 ~ 29999 database errors
	INTERNAL_ERROR int = 20000
)

type EmsError interface {
	Code() int
	Error() string
}

type err struct {
	code int
	msg  string
}

func (e *err) Code() int {
	return e.code
}

func (e *err) Error() string {
	return e.msg
}

func New(code int, msg string) EmsError {
	return &err{
		code: code,
		msg:  msg,
	}
}
