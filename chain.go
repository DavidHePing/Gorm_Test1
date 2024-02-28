package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
)

func chain_test1(db *gorm.DB) {
	var cars = []model.Car{}
	db.Where("price >= ?", 100).Where("type = ?", "Japan").Find(&cars)
	fmt.Println(cars)
}
