package infrastructure

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/infrastructure/database"
	"gorm.io/gorm"
)

type Infrastructure struct {
	DB *gorm.DB
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

func (i *Infrastructure) Build() *Infrastructure {

	i.DB = database.Connect()
	return i
}
