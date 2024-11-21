package users

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Loginname string
	Password  string
}

// 创建用户
func NewUser(id uuid.UUID, loginname, password string) (*User, error) {

	if id == uuid.Nil {
		return nil, errors.New("id cannot be nil")
	}

	if loginname == "" {
		return nil, errors.New("loginname cannot be empty")
	}

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	return &User{
		ID:        id,
		Loginname: loginname,
		Password:  password,
	}, nil
}

// 修改密码
func (u *User) ChangePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	u.Password = password

	return nil
}

// 校验密码
func (u *User) ValidatePassword(password string) bool {
	if u.Password == password {
		return true
	}

	return false
}
