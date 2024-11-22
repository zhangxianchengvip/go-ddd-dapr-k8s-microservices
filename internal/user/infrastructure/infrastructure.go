package infrastructure

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/database"
	"gorm.io/gorm"
)

type Infrastructure struct {
	DB *gorm.DB
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

func (i *Infrastructure) Build() *Infrastructure {

	i.DB = database.PostgresGormConnection("host= localhost user=postgres password=1 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	return i
}
