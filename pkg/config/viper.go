package config

import "github.com/spf13/viper"

func NewViperConfig() *viper.Viper {
	return viper.New()
}
