package main

import (
	"fmt"

	"gorm.io/gorm"
)

func select_test1(db *gorm.DB) {
	var car Car

	//db.Debug().First(&car)//show single sql
	db.First(&car)

	fmt.Println(car)
}
