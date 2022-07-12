package database

import (
	"fmt"
	"project/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)

	// dsn := "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True"
	fmt.Println("\n-----------------------\n--------------------dsn:\n\n", dsn)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn))
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	// sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	// sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	// sqlDB.SetConnMaxLifetime(10*time.Second)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Db.Error != nil {
		fmt.Printf("database error %v", Db.Error)
	}
}

// func init() {
// 	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
// 	// 	utils.DbUser,
// 	// 	utils.DbPassword,
// 	// 	utils.DbHost,
// 	// 	utils.DbPort,
// 	// 	utils.DbName,
// 	// )
// 	Db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True"))
// 	if err != nil {
// 		fmt.Printf("mysql connect error %v", err)
// 	}
// 	if Db.Error != nil {
// 		fmt.Printf("database error %v", Db.Error)
// 	}

// }
