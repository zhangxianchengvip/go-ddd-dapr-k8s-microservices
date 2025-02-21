package apis

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/admin/application/features/users/cmds"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/admin/application/features/users/queries"

	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

// @BasePath /api/v1

// UserCreate 创建新用户
// @Summary 创建新用户
// @Description 创建一个新用户，并返回用户信息
// @Tags users
// @Accept json
// @Produce json
// @Param body body cmds.CreateUserCommand true "用户详细信息"
// @Success 200 {object} cmds.CreateUserCommandResponse "用户创建成功"
// @Failure 400 {object} map[string]string "请求参数错误"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /users [post]
func Create(c *gin.Context) {

	var cmd cmds.CreateUserCommand

	if err := c.Bind(&cmd); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := mediatr.Send[*cmds.CreateUserCommand, *cmds.CreateUserCommandResponse](c.Request.Context(), &cmd)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)

}

// @BasePath /api/v1

// User 查询用户信息
// @Summary 查询用户信息
// @Description 查询一个用户，并返回用户信息
// @Tags users
// @Accept json
// @Produce json
// @Param id query string true "用户ID"
// @Success 200 {object} queries.UserByIDQueryResponse "查询成功"
// @Failure 400 {object} map[string]string "请求参数错误"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /users [get]
func User(c *gin.Context) {

	var query queries.UserByIDQuery

	if err := c.BindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := mediatr.Send[*queries.UserByIDQuery, *queries.UserByIDQueryResponse](c.Request.Context(), &query)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)

}

// @BasePath /api/v1

// User 用户登录
// @Summary 用户登录
// @Description 用户登录，并返回token
// @Tags users
// @Accept json
// @Produce json
// @Param body body queries.UserLoginQuery true "用户登录"
// @Success 200 {object} queries.UserByIDQueryResponse "查询成功"
// @Failure 400 {object} map[string]string "请求参数错误"
// @Failure 500 {object} map[string]string "服务器内部错误"
// @Router /login [post]
func Login(c *gin.Context) {

	var query queries.UserLoginQuery

	if err := c.BindJSON(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := mediatr.Send[*queries.UserLoginQuery, *queries.UserLoginQueryResponse](c.Request.Context(), &query)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jwt(response)
	c.JSON(200, response)
}

func jwt(query *queries.UserLoginQueryResponse) string {
	// TODO: jwt
	return ""
}
