package config

import (
	"fmt"
	"time"
)

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	TTLModel int    `mapstructure:"ttl_model"`
}

func (r *Redis) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

func (r *Redis) Index() int {
	return 0
}

func (r *Redis) ConfigInfo() string {
	return fmt.Sprintf("%+v", r)
}

func (r *Redis) GetTTLModel() time.Duration {
	return time.Duration(r.TTLModel) * time.Minute
}
