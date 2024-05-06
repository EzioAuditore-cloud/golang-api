package database

import (
	"fmt"
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
	dataBytes, err := os.ReadFile("./DB.yaml")
	if err != nil {
		logger.StructLog("Error", "db config init ReadFile err: %v", err)
	}
	config := DBConfig{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		logger.StructLog("Error", "db config init Unmarshal err: %v", err)
	}
	config.Mysql = Mysql{
		Host: os.Getenv("MYSQL_HOST"),
		Port: os.Getenv("MYSQL_PORT"),
	}
	dbUser := os.Getenv("MYSQL_USER")
	dbPwd := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, config.Mysql.Host, config.Mysql.Port, dbName)
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
