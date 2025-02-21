package infrastructure

import (
	"github.com/spf13/viper"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/database"
	"gorm.io/gorm"
)

var (
	database_conn_config = "database.conn"
)

func NewRespository(viper *viper.Viper) *gorm.DB {

	conn := viper.GetString(database_conn_config)

	db, err := database.PostgresGormConnection(conn)

	if err != nil {
		panic(err)
	}

	return db
}
