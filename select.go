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

	map1 := map[string]interface{}{}
	db.Model(&Car{}).First(&map1)
	fmt.Print(map1)
}
