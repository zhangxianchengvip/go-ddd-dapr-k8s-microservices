package web

import (
	"context"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/api/openapi-spec/user"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/cmd/user/app/web/routes"
	"go.uber.org/fx"
)

var (
	port      = ":8081"
	gin_group = "/api/v1/"
	bathPath  = "/api/v1/"
	swagger   = "/swagger/*any"
)

func NewGin() *gin.Engine {
	return gin.Default()
}

func RunServer(lifecycle fx.Lifecycle, r *gin.Engine) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				routes.Map(r.Group(gin_group))
				docs.SwaggerInfo.BasePath = bathPath
				r.GET(swagger, ginSwagger.WrapHandler(swaggerfiles.Handler))
				_ = r.Run(port) // 启动 Gin 服务
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 停止逻辑（如果需要）
			return nil
		},
	})
}

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(NewGin),
		fx.Invoke(RunServer),
	}

}
