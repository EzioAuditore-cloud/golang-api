package config

import (
	"os"
	"project/middleWare/logger"

	"gopkg.in/yaml.v3"
)

type GlobleConfig struct {
	JwtSecret string `yaml:"JwtSecret"`
}

var GlobleConf GlobleConfig

func init() {
	dataBytes, err := os.ReadFile("../config/globleConf.yaml")
	if err != nil {
		logger.StructLog("Error", "Globle config init ReadFile err: %v", err)
	}
	err = yaml.Unmarshal(dataBytes, &GlobleConf)
	if err != nil {
		logger.StructLog("Error", "Globle config init Unmarshal err: %v", err)
	}
	logger.StructLog("Info", "全局 配置成功！")
}
