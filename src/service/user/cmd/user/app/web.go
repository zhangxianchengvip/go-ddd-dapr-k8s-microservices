package app

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/cmd/user/app/routes"
	docs "github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/docs"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/application"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/domain"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/infrastructure"
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
	routes.Map(w.gin.Group("/api/v1/"))
	docs.SwaggerInfo.BasePath = "/api/v1"
	w.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	w.gin.Run(":8080")
}
