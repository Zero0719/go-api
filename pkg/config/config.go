package config

import (
	"os"

	"github.com/Zero0719/go-api/helpers"
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {
	viper = viperlib.New()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("appenv")
	viper.AutomaticEnv()
	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	loadEnv(env)

	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(envSuffix string) {
	envPath := ".env"

	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			envPath = filepath
		}
	}

	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// 读取 env 中的配置信息
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func Get[T any](path string, defaultValue ...interface{}) T {
	if value := internalGet(path, defaultValue...); value != nil {
		return value.(T)
	}

	var fallback T
	return fallback
}

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

func GetStringMapString(path string, defaultValue ...interface{}) map[string]string {
	return viper.GetStringMapString(path)
}