package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Database *Database `mapstructure:"db"`
	Redis    *Redis    `mapstructure:"redis"`
}

var (
	instance *Config
	once     sync.Once
	mutex    sync.Mutex
)

func initEnvironment() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.Set("env", "development")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Failed read file config : %s", err.Error())
	}

	if err := viper.UnmarshalKey(viper.GetString("env"), &instance); err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
}

func GetConfig() *Config {
	once.Do(initEnvironment)
	mutex.Lock()
	defer mutex.Unlock()
	return instance
}
