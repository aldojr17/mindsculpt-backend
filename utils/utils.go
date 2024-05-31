package utils

import (
	"fmt"
	"mindsculpt/config"
)

func GetUrl(url string) string {
	return fmt.Sprintf("%s%s", config.GetConfig().API.URL, url)
}
