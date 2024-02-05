package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
)

func select_count_test1(db *gorm.DB) {
	var count int64
	db.Table("cars").Select("id").Count(&count)
	fmt.Println(count)
}

func select_raw_sql_test1(db *gorm.DB) {
	var cars []model.Car
	db.Raw("select * from cars where id = ? and name = ? and price = ?", 1, "Tesla", 1000).Scan(&cars)
	fmt.Println(cars)
}

func select_limit_offset_test1(db *gorm.DB) {
	var users []model.User

	db.Limit(10).Offset(10).Find(&users)
	fmt.Println(users)
}
