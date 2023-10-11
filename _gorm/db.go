// Package _gorm @Author:冯铁城 [17615007230@163.com] 2023-09-14 09:22:28
package _gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库连接参数
var ip = "127.0.0.1"
var port = "3306"
var username = "root"
var password = "root"
var db = "study"

// InitDB 初始化数据库
func InitDB() *gorm.DB {

	//1.拼接DSN
	dsnSuffix := "charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", username, password, ip, port, db, dsnSuffix)

	//2.打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("database connection error")
	}

	//3.返回数据库连接
	return db
}
