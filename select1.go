package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
)

func select_test1(db *gorm.DB) {
	var car model.Car

	//db.Debug().First(&car)//show single sql
	db.First(&car)

	fmt.Println(car)

	map1 := map[string]interface{}{}
	db.Model(&model.Car{}).First(&map1)
	fmt.Print(map1)

	var cars []model.Car

	db.Find(&cars)
	fmt.Println(cars)
}

func select_with_id(db *gorm.DB) {
	var cars []model.Car
	result := db.Find(&cars, 1)
	fmt.Println(cars)
	fmt.Println(result.Error)
}

func select_where_test1(db *gorm.DB) {
	var cars []model.Car

	result := db.Where("name = ?", "Tesla").Find(&cars)
	fmt.Println(cars)

	result = db.Where("price >= ?", 200).Find(&cars)
	result = db.Where("price >= ?", "200").Find(&cars)
	if result.Error != nil {
		fmt.Println("sql failed", result.Error)
	}

	fmt.Println(cars)
}

func select_where_and_test1(db *gorm.DB) {
	var cars []model.Car

	db.Where("name = ?", "Tesla").Or("price >= ?", 200).Find(&cars)
	fmt.Println(cars)

	db.Where("name = ? AND price >= ?", "Tesla", 200).Find(&cars)
	fmt.Println(cars)
}
