package infrastructure

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/database"
	"gorm.io/gorm"
)

func NewRespository() *gorm.DB {

	db, err := database.PostgresGormConnection("host= localhost user=postgres password=1 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai")

	if err != nil {
		panic(err)
	}

	return db
}
