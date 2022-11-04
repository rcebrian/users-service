package users

import "errors"

var ErrInvalidUserID = errors.New("invalid User ID")

var ErrEmptyUserName = errors.New("the field User Name can not be empty")

var ErrEmptyFirstname = errors.New("the field Firstname can not be empty")

var ErrNotFound = errors.New("user not found")
