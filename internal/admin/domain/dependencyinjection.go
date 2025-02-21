package domain

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/admin/domain/users"
	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(users.NewManager),
	}

}
