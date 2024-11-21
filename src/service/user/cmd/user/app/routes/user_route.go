package routes

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/cmd/user/app/apis"

	"github.com/gin-gonic/gin"
)

func MapUserRoutes(r *gin.RouterGroup) {
	r.POST("/users", apis.Create)
	r.GET("/users", apis.User)
}
