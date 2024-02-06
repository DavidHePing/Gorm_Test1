package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
)

func select_preload_test1(db *gorm.DB) {
	var users []model.User

	result := db.Preload("Tasks").Omit("random_number").Where("id in ?", []int{1, 2}).Find(&users)
	if result.Error != nil {
		fmt.Println("select error!!!", result.Error)
	}

	fmt.Println(users)
	// fmt.Println(result)
}
