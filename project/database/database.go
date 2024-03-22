package database

import (
	"os"
	"project/middleWare/logger"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

type DBConfig struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type Mysql struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func init() {
	dataBytes, err := os.ReadFile("../database/config/DB.yaml")
	if err != nil {
		logger.StructLog("Error", "db config init ReadFile err: %v", err)
	}
	config := DBConfig{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		logger.StructLog("Error", "db config init Unmarshal err: %v", err)
	}

	dsn := "root:123456@tcp(" + config.Mysql.Host + ":" + config.Mysql.Port + ")/test?charset=utf8&parseTime=True"
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.StructLog("Error", "db config init mysql connect err: %v", err)
		panic(err)
	}
	if Db.Error != nil {
		logger.StructLog("Error", "db config init database err: %v", Db.Error)
	}
	logger.StructLog("Info", "DB 配置成功！")
}
