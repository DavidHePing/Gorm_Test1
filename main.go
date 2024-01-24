package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Car struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

func main() {
	connect_string := "host=localhost user=admin password=1234 dbname=testDb port=5432"
	db, err := gorm.Open(postgres.Open(connect_string), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed!!!!!")
	}

	var car Car

	db.First(&car)

	fmt.Println("1234")
	fmt.Println(car)
}
