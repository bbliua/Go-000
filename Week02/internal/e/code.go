package e

import (
	"fmt"
)

type ecode struct {
	code int
	msg  string
}

func (e *ecode) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.code, e.msg)
}

func New(code int, msg string) error {
	e := &ecode{
		code: code,
		msg:  msg,
	}
	return e
}

// ErrRecordNotFound represents for no record error
var ErrRecordNotFound = New(404, "Record Not Found")

// ErrQueryFail describes action query fail
var ErrQueryFail = New(500, "Query Fail")
