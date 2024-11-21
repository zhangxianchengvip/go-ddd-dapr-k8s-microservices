package routes

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/cmd/user/app/apis"

	"github.com/gin-gonic/gin"
)

func MapUserRoutes(r *gin.RouterGroup) {

	g := r.Group("users")
	g.POST("", apis.Create)
	g.GET("", apis.User)
}