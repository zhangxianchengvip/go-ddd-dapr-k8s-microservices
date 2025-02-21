package cmds

import (
	"context"

	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/admin/domain/users"
	"gorm.io/gorm"
)

type CreateUserCommand struct {
	Loginname string `json:"loginname"` //用户名
	Password  string `json:"password"`  //密码
}

type CreateUserCommandResponse struct {
	Success bool `json:"success"`
}

type CreateUserCommandHandler struct {
	repository *gorm.DB
	manager    *users.UserManager
}

func NewCreateUserCommandHandler(repository *gorm.DB, manager *users.UserManager) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		repository: repository,
		manager:    manager,
	}
}

func (h *CreateUserCommandHandler) Handle(ctx context.Context, command *CreateUserCommand) (*CreateUserCommandResponse, error) {

	user, err := h.manager.Create(command.Loginname, command.Password)

	if err != nil {
		return nil, err
	}

	result := h.repository.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &CreateUserCommandResponse{Success: true}, nil
}
