package database

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/domain/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresGormConnection(conn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&users.User{})

	return db
}
