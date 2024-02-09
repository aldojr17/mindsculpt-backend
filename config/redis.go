package config

import "fmt"

type Redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
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
