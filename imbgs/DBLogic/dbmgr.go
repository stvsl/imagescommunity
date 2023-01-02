package Database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB // 全局连接池

func Connect() {
	if _db == nil {
		dsn := "root:stvsl2060@tcp(120.48.17.34:27621)/imagecommunity?charset=utf8mb4&parseTime=True&loc=Local"
		_db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		fmt.Println("数据库连接成功")
		SetPool()
	} else {
		ReConnect()
	}
}

func ReConnect() {
	Close()
	Connect()
}

func Close() {
	if _db != nil {
		sqldb, _ := _db.DB()
		sqldb.Close()
	}
	_db = nil
}

func SetPool() {
	dbc, err := _db.DB()
	if err != nil {
		fmt.Println("SetPoolerr" + err.Error())
	}
	dbc.SetMaxIdleConns(10)
	dbc.SetMaxOpenConns(30)
	dbc.SetConnMaxLifetime(time.Duration(2 * int(time.Hour)))
	fmt.Println("连接池设置成功")
}

func GetConn() *gorm.DB {
	if _db == nil {
		fmt.Println("GetConnErr db is nil")
	}
	return _db
}
