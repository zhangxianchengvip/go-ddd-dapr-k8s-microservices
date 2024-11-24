package infrastructure

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/database"
	"gorm.io/gorm"
)

type Infrastructure struct {
	db_conn string
	DB      *gorm.DB
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		db_conn: "host= localhost user=postgres password=1 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai",
	}
}

func (i *Infrastructure) Build() *Infrastructure {

	db, err := database.PostgresGormConnection(i.db_conn)

	if err != nil {
		panic(err)
	}

	i.DB = db

	return i
}
