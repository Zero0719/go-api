package utils

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Fatal().Msgf("读取配置文件失败: %v", err)
	}
	Config = viper.GetViper()
}