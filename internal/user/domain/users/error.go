package users

import "errors"

var (
	ErrRepositoryNil     = errors.New("repository is nil")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrIDEmpty           = errors.New("id cannot be empty")
	ErrLoinnameEmpty     = errors.New("loginname cannot be empty")
	ErrPasswordEmpty     = errors.New("password cannot be empty")
)
