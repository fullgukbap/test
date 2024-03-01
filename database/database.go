package database

import (
	"simple-user/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	// 데이터베이스 인스턴스 생성
	dsn := "root:Qwe123!@#@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	// 데이터베이스 마이그레이션
	err = db.AutoMigrate(new(entity.User))
	if err != nil {
		panic(err)
	}

	// DB.Create(&entity.User{Username: "guja", Password: "guja"})
}
