package cqrs

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/application/users/cmds"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/application/users/queries"

	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/domain/users"

	"github.com/mehdihadeli/go-mediatr"
	"gorm.io/gorm"
)

func User(db *gorm.DB) {

	manager, err := users.NewManager(db)

	if err != nil {
		panic(err)
	}

	merr := mediatr.RegisterRequestHandler[*cmds.CreateUserCommand, *cmds.CreateUserCommandResponse](cmds.NewCreateUserCommandHandler(db, manager))
	if merr != nil {
		panic(merr)
	}

	merr = mediatr.RegisterRequestHandler[*queries.UserByIDQuery, *queries.UserByIDQueryResponse](queries.NewUserByIdQueryHandler(db))

	if merr != nil {
		panic(merr)
	}
}
