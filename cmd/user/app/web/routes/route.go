package routes

import (
	"github.com/gin-gonic/gin"
)

func Map(r *gin.RouterGroup) {
	MapUserRoutes(r)
}
