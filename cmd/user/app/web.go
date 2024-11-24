package app

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/api/openapi-spec/user"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/cmd/user/app/routes"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/domain"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/infrastructure"
)

var (
	port      = ":8081"
	gin_group = "/api/v1/"
	bathPath  = "/api/v1/"
	swagger   = "/swagger/*any"
)

type webApp struct {
	gin            *gin.Engine
	domain         *domain.Domain
	application    *application.Application
	infrastructure *infrastructure.Infrastructure
}

type webApplicatonBuilder struct {
}

func newWebApplicatonBuilder() *webApplicatonBuilder {
	return &webApplicatonBuilder{}
}

func (w *webApplicatonBuilder) Build() *webApp {

	i := infrastructure.NewInfrastructure().Build()

	a := application.NewApplication(i.DB).Build()

	d := domain.NewDomain().Build()

	g := gin.Default()

	return &webApp{
		gin:            g,
		domain:         d,
		application:    a,
		infrastructure: i,
	}
}

func (w *webApp) Run() {
	routes.Map(w.gin.Group(gin_group))
	docs.SwaggerInfo.BasePath = bathPath
	w.gin.GET(swagger, ginSwagger.WrapHandler(swaggerfiles.Handler))
	w.gin.Run(port)
}
