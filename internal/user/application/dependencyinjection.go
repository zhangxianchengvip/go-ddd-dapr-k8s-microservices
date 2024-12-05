package application

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/cqrs"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/cmds"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/queries"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/config"
	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(cmds.NewCreateUserCommandHandler),
		fx.Provide(queries.NewUserByIdQueryHandler),
		fx.Provide(queries.NewUserLoginQueryHandler),
		fx.Provide(cqrs.NewUserCQRS),
		fx.Provide(config.NewViperConfig),
	}
}
