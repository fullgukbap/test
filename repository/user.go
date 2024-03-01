package repository

import (
	"simple-user/database"
	"simple-user/entity"
)

func Create(username string, password string) error {
	return database.DB.Create(&entity.User{Username: username, Password: password}).Error
}

func Read(username string) (*entity.User, error) {
	user := new(entity.User)
	err := database.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

// func Update(ID uint, username string, password string) error {
// 	user := &entity.User{}
// 	database.DB.Model(&entity.User{}).Updates(entity.User{Price: 200, Code: "F42"})
// }

// func Delete(ID uint) error {

// }
