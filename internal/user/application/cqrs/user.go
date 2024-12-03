package cqrs

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/cmds"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/queries"

	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/domain/users"

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

	merr = mediatr.RegisterRequestHandler[*queries.UserLoginQuery, *queries.UserLoginQueryResponse](queries.NewUserLoginQueryHandler(db))

	if merr != nil {
		panic(merr)
	}
}
