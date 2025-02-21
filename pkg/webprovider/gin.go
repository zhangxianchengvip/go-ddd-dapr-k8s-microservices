package webprovider

import "github.com/gin-gonic/gin"

func NewGin() *gin.Engine {
	return gin.Default()
}
