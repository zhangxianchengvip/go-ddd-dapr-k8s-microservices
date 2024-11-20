package cmds

import (
	"context"

	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/domain/users"
	"gorm.io/gorm"
)

type CreateUserCommand struct {
	loginname string
	password  string
}

type CreateUserCommandResponse struct {
	Success bool
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

	user, err := h.manager.Create(command.loginname, command.password)

	if err != nil {
		return nil, err
	}

	result := h.repository.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &CreateUserCommandResponse{Success: true}, nil
}
