package apis

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/user/internal/application/users/cmds"

	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

func UserCreate(c *gin.Context) {

	var cmd cmds.CreateUserCommand

	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := mediatr.Send[cmds.CreateUserCommand, cmds.CreateUserCommandResponse](c.Request.Context(), cmd)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)

}
