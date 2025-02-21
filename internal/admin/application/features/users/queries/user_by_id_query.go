package queries

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/admin/domain/users"
	"gorm.io/gorm"
)

type UserByIDQuery struct {
	ID string `json:"id"  form:"id"` // 用户id
}

type UserByIDQueryResponse struct {
	ID        uuid.UUID `json:"id"`        // 用户id
	Loginname string    `json:"loginname"` // 用户登录名
}

type UserByIdQueryHandler struct {
	repo *gorm.DB
}

func NewUserByIdQueryHandler(repo *gorm.DB) *UserByIdQueryHandler {
	return &UserByIdQueryHandler{
		repo: repo,
	}
}

func (h *UserByIdQueryHandler) Handle(ctx context.Context, query *UserByIDQuery) (*UserByIDQueryResponse, error) {

	var user users.User

	id, err := uuid.Parse(query.ID)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	if err := h.repo.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &UserByIDQueryResponse{
		ID:        user.ID,
		Loginname: user.Loginname,
	}, nil
}
