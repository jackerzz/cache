package gdb

import (
	"cache/setting"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Setup() (err error) {
	user := setting.DatabaseSetting.User
	password := setting.DatabaseSetting.Password
	host := setting.DatabaseSetting.Host
	dbName := setting.DatabaseSetting.Name
	dbType := setting.DatabaseSetting.Type
	DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
			fmt.Println("DB conn error:",err)
	}
	fmt.Println("Database connection successful")
	DB.SingularTable(true)
	// Waiting for model data Migration
	DB.AutoMigrate()
	return nil
}

func CloseDB() {
	defer DB.Close()
}