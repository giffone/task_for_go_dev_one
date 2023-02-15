package model

import "errors"

var (
	ErrTimeOut   = errors.New("time out")
	ErrUserExist = errors.New("user already exist")
)
