package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Car struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

func main() {
	connect_string := "host=localhost user=admin password=1234 dbname=testDb port=5432"
	db, err := gorm.Open(postgres.Open(connect_string), &gorm.Config{
		//show all sql
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Failed!!!!!")
	}

	var car Car

	db.First(&car)
	//db.Debug().First(&car)//show single sql

	fmt.Println(car)
}

func InsertTest() {

}
