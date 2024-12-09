package web

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/api/openapi-spec/user"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/cmd/user/app/web/routes"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/cqrs"
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

// Configuretion 配置
func Configuretion(v *viper.Viper) {
	v.Set("database.conn", "host=192.168.110.21 user=postgres password=1 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai")
}

// RunServer 启动服务
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
	}

}

func Invoke() []fx.Option {
	return []fx.Option{
		fx.Invoke(Configuretion),

		fx.Invoke(func(uc *cqrs.UserCQRS) {
			uc.User()
		}),

		fx.Invoke(RunServer),
	}
}
