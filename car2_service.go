package main

import (
	"Gorm_Test1/model"

	"gorm.io/gorm"
)

func test_car2(db *gorm.DB) {
	c := model.Car2{Id: 1, IsDeleted: true}
	// c.Select(db)
	c.Delete(db)
	c.SelectAll(db)
}
