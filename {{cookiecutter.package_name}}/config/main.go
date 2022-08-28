package config

import (
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"os"

	"log"
)

type Config interface {
	GetInt(key string) int
	GetString(key string) string
	Get(key string) interface{}
}

type viperConfig struct {
	viper *viper.Viper
}

func (v viperConfig) GetInt(key string) int {
	return v.viper.GetInt(key)
}

func (v viperConfig) Get(key string) interface{} {
	return v.viper.Get(key)
}

func (v viperConfig) GetString(key string) string {
	return v.viper.GetString(key)
}

var config *viperConfig

func init() {
	config = nil
}

func setDefault(viperObj *viper.Viper) {
	viperObj.SetDefault("server_port", 8000)
}

func createConfig() *viperConfig {
	newViper := viper.New()

	configName := os.Getenv("CONFIG_FILENAME")
	if configName == "" {
		configName = "local"
	}

	newViper.SetConfigName(configName)
	newViper.AddConfigPath(".")
	newViper.SetConfigType("yaml")

	// Env
	newViper.SetEnvPrefix("conf")
	newViper.AutomaticEnv()

	setDefault(newViper)

	if err := newViper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return &viperConfig{viper: newViper}
}

func GetConfig() Config {
	if config == nil {
		config = createConfig()
	}

	return config
}

func Provide(container *dig.Container) {
	container.Provide(GetConfig)
}
