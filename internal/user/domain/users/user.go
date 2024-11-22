package users

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`                 // 主键
	Loginname string    `gorm:"type:varchar(255);unique_index" json:"loginname"` // 登录名
	Password  string    `gorm:"type:varchar(255)" json:"password"`               // 密码
	Realname  string    `gorm:"type:varchar(255)" json:"realname"`               // 真实姓名
	Email     string    `gorm:"type:varchar(255)" json:"email"`                  // 邮箱
	Phone     string    `gorm:"type:varchar(255)" json:"phone"`                  // 手机号
	RoleId    uuid.UUID `gorm:"type:uuid" json:"role_id"`                        // 角色ID
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
