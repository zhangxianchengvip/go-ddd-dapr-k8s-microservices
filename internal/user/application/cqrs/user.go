package cqrs

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/cmds"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/queries"

	"github.com/mehdihadeli/go-mediatr"
	"gorm.io/gorm"
)

func User(
	db *gorm.DB,
	createUserHandler *cmds.CreateUserCommandHandler,
	userByIdQueryHandler *queries.UserByIdQueryHandler,
	userLoginQueryHandler *queries.UserLoginQueryHandler,

) {

	merr := mediatr.RegisterRequestHandler[*cmds.CreateUserCommand, *cmds.CreateUserCommandResponse](createUserHandler)
	if merr != nil {
		panic(merr)
	}

	merr = mediatr.RegisterRequestHandler[*queries.UserByIDQuery, *queries.UserByIDQueryResponse](userByIdQueryHandler)

	if merr != nil {
		panic(merr)
	}

	merr = mediatr.RegisterRequestHandler[*queries.UserLoginQuery, *queries.UserLoginQueryResponse](userLoginQueryHandler)

	if merr != nil {
		panic(merr)
	}
}
