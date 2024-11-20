package routes

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/cmd/user/webapp/apis"

	"github.com/gin-gonic/gin"
)

func MapUserRoutes(r *gin.RouterGroup) {

	g := r.Group("user")
	g.POST("", apis.UserCreate)
}
