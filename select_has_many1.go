package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func select_many_test1(db *gorm.DB) {
	var users []model.User
	result := db.Find(&users)
	fmt.Println(result)
}

func select_many_test2(db *gorm.DB) {
	var users []model.User
	// db.DryRun = true
	db.Preload(clause.Associations).Omit("random_number").Find(&users)
	fmt.Println(users)
}
