package database

import (
	"errors"

	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/domain/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresGormConnection(conn string) (*gorm.DB, error) {

	if conn == "" {
		return nil, errors.New("empty connection string")
	}

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&users.User{})

	return db, nil
}
