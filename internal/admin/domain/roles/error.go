package roles

import "errors"

var (
	ErrNameEmpty = errors.New("name cannot be empty")
	ErrCodeEmpty = errors.New("code cannot be empty")
)
