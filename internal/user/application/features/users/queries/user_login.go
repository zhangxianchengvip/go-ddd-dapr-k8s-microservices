package queries

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/domain/users"
	"gorm.io/gorm"
)

type UserLoginQuery struct {
	Loginname string `json:"loginname" binding:"required"` //用户名
	Password  string `json:"password" binding:"required"`  //密码
}

type UserLoginQueryResponse struct {
	Id        uuid.UUID
	Loginname string
}

type UserLoginQueryHandler struct {
	db *gorm.DB
}

func NewUserLoginQueryHandler(db *gorm.DB) *UserLoginQueryHandler {
	return &UserLoginQueryHandler{db: db}
}

func (h *UserLoginQueryHandler) Handle(ctx context.Context, query *UserLoginQuery) (*UserLoginQueryResponse, error) {

	var user users.User

	if err := h.db.Where("loginname = ? ", query.Loginname).First(&user).Error; err != nil {

		return nil, err
	}

	if !user.ValidatePassword(query.Password) {
		return nil, errors.New("password is not correct")
	}

	return &UserLoginQueryResponse{
		Id:        user.ID,
		Loginname: user.Loginname,
	}, nil
}
