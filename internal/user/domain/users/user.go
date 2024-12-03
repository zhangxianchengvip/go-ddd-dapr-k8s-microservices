package users

import (
	"github.com/google/uuid"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/ddd"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/tools"
)

type User struct {
	ddd.AggregateRoot[uuid.UUID]           // 聚合根
	Loginname                    string    // 登录名
	Password                     string    // 密码
	Realname                     string    // 真实姓名
	Email                        string    // 邮箱
	Phone                        string    // 手机号
	RoleId                       uuid.UUID // 角色ID
}

// 创建用户
func NewUser(id uuid.UUID, loginname, password string) (*User, error) {

	if id == uuid.Nil {
		return nil, ErrIDEmpty
	}

	if tools.StringIsEmplyOrWhiteSpace(loginname) {
		return nil, ErrLoinnameEmpty
	}

	if tools.StringIsEmplyOrWhiteSpace(password) {
		return nil, ErrPasswordEmpty
	}

	return &User{
		AggregateRoot: ddd.NewAggregateRoot(id),
		Loginname:     loginname,
		Password:      password,
	}, nil
}

// 修改密码
func (u *User) ChangePassword(password string) error {
	if tools.StringIsEmplyOrWhiteSpace(password) {
		return ErrPasswordEmpty
	}

	u.Password = password

	// 发布密码修改事件
	event := PasswordChangedEvent{
		DomainEvent: ddd.NewDomainEvent(),
		UserID:      u.ID.String(),
	}

	u.AddDomainEvent(event)

	return nil
}

// 校验密码
func (u *User) ValidatePassword(password string) bool {
	if u.Password == password {
		return true
	}

	return false
}
