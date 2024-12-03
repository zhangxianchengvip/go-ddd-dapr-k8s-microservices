package application

import (
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/internal/user/application/cqrs"
	"gorm.io/gorm"
)

func NewApplication(db *gorm.DB) *Application {
	return &Application{db: db}
}

type Application struct {
	db *gorm.DB
}

func (a *Application) Build() *Application {

	// load CQRS
	{
		cqrs.User(a.db)
	}
	return a
}
