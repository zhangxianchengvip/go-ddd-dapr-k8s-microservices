package application

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/cqrs"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/cmds"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/queries"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(cmds.NewCreateUserCommandHandler),
		fx.Provide(queries.NewUserByIdQueryHandler),
		fx.Provide(queries.NewUserLoginQueryHandler),

		fx.Invoke(func(
			g *gorm.DB,
			c *cmds.CreateUserCommandHandler,
			q *queries.UserByIdQueryHandler,
			l *queries.UserLoginQueryHandler) {
			cqrs.User(g, c, q, l)
		}),
	}
}
