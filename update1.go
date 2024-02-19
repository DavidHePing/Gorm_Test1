package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
)

func save_test1(db *gorm.DB) {
	car := model.Car{Id: 1, Price: 2000, Name: "Tesla"}
	db.Save(&car)
}

func udpate_test1(db *gorm.DB) {
	car := model.Car{Id: 1, Price: 2000}
	db.Model(&car).Update("price", 1000)
	result := db.Model(&car).Select("price").Updates(car)
	fmt.Println("Row affected:", result.RowsAffected)

	var carFromDb model.Car
	db.Find(&carFromDb, 1)
	fmt.Println(carFromDb)
}
