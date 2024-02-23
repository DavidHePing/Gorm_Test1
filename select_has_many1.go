package main

import (
	"Gorm_Test1/model"

	"gorm.io/gorm"
)

func select_many_test1(db *gorm.DB) {
	var users []model.User
	result := db.Find(&users)
	print(result)
}

func select_many_test2(db *gorm.DB) {
	var users []model.User
	db.Joins("User").Joins("Task").Find(&users)
	print(users)
}
