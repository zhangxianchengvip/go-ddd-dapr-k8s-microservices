package infrastructure

import (
	"github.com/spf13/viper"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/database"
	"gorm.io/gorm"
)

func NewRespository(viper *viper.Viper) *gorm.DB {

	conn := viper.GetString("database.conn")

	db, err := database.PostgresGormConnection(conn)

	if err != nil {
		panic(err)
	}

	return db
}
