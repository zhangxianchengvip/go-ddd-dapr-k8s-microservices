package cqrs

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/application/users/cmds"

	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/domain/users"

	"github.com/mehdihadeli/go-mediatr"
	"gorm.io/gorm"
)

func User(db *gorm.DB) {

	manager, err := users.NewManager(db)

	if err != nil {
		panic(err)
	}

	handler := cmds.NewCreateUserCommandHandler(db, manager)

	mediatr.RegisterRequestHandler[*cmds.CreateUserCommand, *cmds.CreateUserCommandResponse](handler)
}
