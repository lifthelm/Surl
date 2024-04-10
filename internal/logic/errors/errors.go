package errors

import (
	"errors"
)

var (
	ErrCantGetDBConnectionString = errors.New("cant get db connection string")
	ErrCantAddNewLinkInProject   = errors.New("cant add new links in project")
	ErrNoRoute                   = errors.New("can match any route")
	ErrCantConvertToLogic        = errors.New("cant convert data to logic model")
	ErrUnknownUserRole           = errors.New("cant determine user role")
	ErrFilterAllGeneral          = errors.New("cant user all general filter")
	ErrCantFind                  = errors.New("cant find record")
	ErrAuth                      = errors.New("auth failed")
	ErrCantUpdate                = errors.New("cant update data")
	ErrCantRequest               = errors.New("cant request data")
	ErrCantInsert                = errors.New("cant insert record")
	ErrAlreadyExist              = errors.New("record already exist")
)
