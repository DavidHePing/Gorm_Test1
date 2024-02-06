package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
)

func select_join_test1(db *gorm.DB) {
	var users []model.User

	result := db.Joins("join tasks on tasks.user_id = test_users.id").Select("test_users.id", "name", "age").Find(&users)
	if result.Error != nil {
		fmt.Println("select error!!!", result.Error)
	}

	// fmt.Println(users)
	fmt.Println(result.Statement)
}
