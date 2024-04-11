package errors

import "errors"

var (
	ErrWrongChoice      = errors.New("wrong option number")
	ErrWrongInput       = errors.New("wrong input")
	ErrHandlerExecution = errors.New("wrong handler execution")
	ErrServiceInit      = errors.New("cant initialize service")
	ErrUnknownRole      = errors.New("unknown role")
	ErrUnknownPlatform  = errors.New("unknown platform")
	ErrUnknownGeo       = errors.New("unknown geo")
)
