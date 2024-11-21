package database

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/domain/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dsn := "host=localhost user=postgres password=1 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&users.User{})

	return db
}
