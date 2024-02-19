package main

import (
	"Gorm_Test1/model"

	"gorm.io/gorm"
)

func save_test1(db *gorm.DB) {
	car := model.Car{Id: 1, Price: 2000, Name: "Tesla"}
	db.Save(&car)
}

func udpate_test1(db *gorm.DB) {
	car := model.Car{Id: 1, Price: 2000}
	db.Model(&car).Update("price", 1000)
	db.Model(&car).Select("price").Updates(car)
}
