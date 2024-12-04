package application

import (
	"github.com/spf13/viper"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/cqrs"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/cmds"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/features/users/queries"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/config"
	"go.uber.org/fx"
)

func Configuretion(v *viper.Viper) {
	v.Set("database.conn", "host= localhost user=postgres password=1 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai")
}

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(cmds.NewCreateUserCommandHandler),
		fx.Provide(queries.NewUserByIdQueryHandler),
		fx.Provide(queries.NewUserLoginQueryHandler),
		fx.Provide(cqrs.NewUserCQRS),
		fx.Provide(config.NewViperConfig),

		fx.Invoke(Configuretion),
		fx.Invoke(func(uc *cqrs.UserCQRS) {
			uc.User()
		}),
	}
}
