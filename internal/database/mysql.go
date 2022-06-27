package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConn() *gorm.DB {
	dsn := "root:qwerty@1@tcp(127.0.0.1:3306)/db_rest_api?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	return db
}
