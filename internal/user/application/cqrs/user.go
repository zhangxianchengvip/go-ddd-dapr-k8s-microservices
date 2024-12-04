package cqrs

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/cmds"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/queries"

	"github.com/mehdihadeli/go-mediatr"
)

type UserCQRS struct {
	createUserHandler     *cmds.CreateUserCommandHandler
	userByIdQueryHandler  *queries.UserByIdQueryHandler
	userLoginQueryHandler *queries.UserLoginQueryHandler
}

func NewUserCQRS(
	createUserHandler *cmds.CreateUserCommandHandler,
	userByIdQueryHandler *queries.UserByIdQueryHandler,
	userLoginQueryHandler *queries.UserLoginQueryHandler) *UserCQRS {

	return &UserCQRS{
		createUserHandler:     createUserHandler,
		userByIdQueryHandler:  userByIdQueryHandler,
		userLoginQueryHandler: userLoginQueryHandler,
	}
}

func (uc *UserCQRS) User() {
	merr := mediatr.RegisterRequestHandler[*cmds.CreateUserCommand, *cmds.CreateUserCommandResponse](uc.createUserHandler)
	if merr != nil {
		panic(merr)
	}

	merr = mediatr.RegisterRequestHandler[*queries.UserByIDQuery, *queries.UserByIDQueryResponse](uc.userByIdQueryHandler)

	if merr != nil {
		panic(merr)
	}

	merr = mediatr.RegisterRequestHandler[*queries.UserLoginQuery, *queries.UserLoginQueryResponse](uc.userLoginQueryHandler)

	if merr != nil {
		panic(merr)
	}
}
