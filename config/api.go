package config

import "fmt"

type API struct {
	URL    string `mapstructure:"url"`
	Key    string `mapstructure:"key"`
	Secret string `mapstructure:"secret"`
}

func (a *API) GetHeaderKey() string {
	return fmt.Sprintf("Key %s", a.Key)
}

func (a *API) GetHeaderSecret() string {
	return fmt.Sprintf("Secret %s", a.Secret)
}
