package main

import (
	"flag"
	"os"

	"workWechat_api/pkg/config"
	"workWechat_api/pkg/logging"
	"workWechat_api/web"
)

var confFile string

func main() {
	flag.StringVar(&confFile, "c", "./config.yml", "配置文件路径,默认为./config.yml")
	flag.Parse()
	_, err := os.Stat(confFile)
	if err != nil{
		logging.Panic("arg error: %T",err)
	}
	config.Init(confFile)
	web.Run()

}
