package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
)

func raw_select_test1(db *gorm.DB) {
	var cars = []model.Car2{}
	db.Raw("select * from test_cars where is_deleted is null").Scan(&cars)
	fmt.Println(cars)

	var car = model.Car2{}
	db.Raw("select * from test_cars where is_deleted is null").Scan(&car)
	fmt.Println(car)
}
