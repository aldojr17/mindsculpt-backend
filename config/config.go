package config

import (
	log "mindsculpt/logger"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Database *Database `mapstructure:"db"`
	Redis    *Redis    `mapstructure:"redis"`
	API      *API      `mapstructure:"api"`
}

var (
	instance *Config
	once     sync.Once
	mutex    sync.Mutex
)

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.Set("env", "development")

	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("failed read file config : %s", err.Error())
		panic("Failed read file config")
	}

	if err := viper.UnmarshalKey(viper.GetString("env"), &instance); err != nil {
		log.Errorf("unable to decode into config struct, %v", err)
		panic("Unable to decode")
	}
}

func GetConfig() *Config {
	once.Do(initConfig)
	mutex.Lock()
	defer mutex.Unlock()
	return instance
}
