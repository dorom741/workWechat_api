package main

import (
	"fmt"
	"testing"

	"qiyewx_api/pkg/config"
	"qiyewx_api/pkg/storage"
)

func Test_redis(t *testing.T) {
	config.Init("../config.yml")
	var key = "unitTest"

	if storage.RedisPool().Set(key,"ffff"){
		fmt.Println(storage.RedisPool().Get(key))
	}
	result,err := storage.RedisPool().Get("ffgf")
	fmt.Println("get with not set key err")
	fmt.Println("result: ",result,"err: ",err)
}

func Test_redis_Singleton(t *testing.T) {
	fmt.Println("1:",*storage.RedisPool())
	fmt.Println("2:",*storage.RedisPool())
}
