package main

import (
	"fmt"
	"mindsculpt/config"
)

func main() {
	config := config.GetConfig()
	fmt.Printf("%+v", config.Database)
}
