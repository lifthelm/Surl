package errors

import "errors"

var (
	ErrCantMatchRecord = errors.New("cant match record")
	ErrNoRecordFound   = errors.New("no records found")
)
