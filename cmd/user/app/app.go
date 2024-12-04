package app

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/cmd/user/app/web"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/domain"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/infrastructure"
	"go.uber.org/fx"
)

func Run() {

	domain_di := domain.DependencyInjection()
	application_di := application.DependencyInjection()
	infrastructure_di := infrastructure.DependencyInjection()
	web_di := web.DependencyInjection()

	module := []fx.Option{}
	module = append(module, domain_di...)
	module = append(module, application_di...)
	module = append(module, infrastructure_di...)
	module = append(module, web_di...)

	// 依赖注入
	app := fx.New(
		module...,
	)

	app.Run()
}
