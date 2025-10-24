package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig `yaml:"app"`
	Auth AuthConfig `yaml:"auth"`
	DB DBConfig `yaml:"db"`
	Redis RedisConfig `yaml:"redis"`
}

type AppConfig struct {
	Port int `yaml:"port"`
	Env string `yaml:"env"`
	Salt string `yaml:"salt"`
}

type AuthConfig struct {
	JwtSecret string `yaml:"jwtSecret"`
	JwtExp int `yaml:"jwtExp"`
	JwtRefreshSecret string `yaml:"jwtRefreshSecret"`
	JwtRefreshExp int `yaml:"jwtRefreshExp"`
}

type DBConfig struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Password string `yaml:"password"`
	Database int `yaml:"database"`
	PoolSize int `yaml:"poolSize"`
}

var globalConfig *Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	globalConfig = &Config{}
	err = viper.Unmarshal(globalConfig)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
}

func Get() *Config {
	return globalConfig
}