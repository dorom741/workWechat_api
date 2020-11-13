package main

import (
	"fmt"
	"testing"

	"workWechat_api/pkg/config"
)

func TestConfig(t *testing.T) {
	config.Init("./config.yml")
	fmt.Println("redis:", config.Redis)
	fmt.Println("general", config.General)
	fmt.Println("qiyeInfo:", config.QiyeInfo)
}