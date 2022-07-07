package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}
}
