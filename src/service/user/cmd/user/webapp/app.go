package webapp

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/cmd/user/webapp/routes"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/application/users/cqrs"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func Run() {

	// connect to database
	db := database.Connect()

	// register CQRS
	cqrs.User(db)

	// register routes
	g := gin.Default()
	{
		basePath := g.Group("/api/v1/")

		routes.Map(basePath)

		g.Run(":8080")
	}
}
