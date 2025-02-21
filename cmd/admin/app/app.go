package app

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/cmd/admin/app/web"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/admin/application"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/admin/domain"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/admin/infrastructure"
	"go.uber.org/fx"
)

func Run() {

	domain_di := domain.DependencyInjection()
	infrastructure_di := infrastructure.DependencyInjection()
	application_di := application.DependencyInjection()
	web_di := web.DependencyInjection()
	invoke := web.Invoke()

	module := []fx.Option{}
	module = append(module, domain_di...)
	module = append(module, application_di...)
	module = append(module, infrastructure_di...)
	module = append(module, web_di...)
	module = append(module, invoke...)

	// 依赖注入
	app := fx.New(
		module...,
	)

	app.Run()
}
