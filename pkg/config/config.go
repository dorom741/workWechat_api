package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"workWechat_api/pkg/logging"
)

var (
	Redis = redis{Network : "tcp",Server: "127.0.0.1:6379",DB: 0}
	General  = general{Listen: ":8080",Ssl: false,CertPath: "",KeyPath: ""}
	QiyeInfo = qiyeInfo{}
)


type conf struct {
	General *general
	Redis *redis
	QiyeInfo *qiyeInfo
}

type general struct {
	Listen   string
	Ssl      bool
	LogLevel string
	CertPath string
	KeyPath  string
}


type redis struct {
	Network  string
	Server   string
	Password string
	DB       int
}

type qiyeInfo struct {
	Corpid string
	AppList map[string]string
}

func Init(filePath string) {
	data, _ := ioutil.ReadFile(filePath)
	config := conf{General: &General,Redis: &Redis,QiyeInfo: &QiyeInfo}
	err := yaml.Unmarshal(data, &config)
	logging.SetLevel(General.LogLevel)
	logging.Info("parse config file successfully，loglevel: %s",General.LogLevel)
	if err != nil {
		logging.Error("parse config file error: %T",err)
	}
}




