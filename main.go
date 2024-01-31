package main

import (
	"fmt"
	"strconv"
	"time"

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
		//not do only show sql
		// DryRun: true,
	})

	if err != nil {
		fmt.Println("Failed!!!!!")
	}

	// select_test1(db)
	insert_test1(db)
}

func select_test1(db *gorm.DB) {
	var car Car

	//db.Debug().First(&car)//show single sql
	db.First(&car)

	fmt.Println(car)
}

func insert_test1(db *gorm.DB) {
	users := []*User{}

	for i := 0; i < 30; i++ {
		users = append(users, &User{
			Name:     strconv.Itoa(i),
			Age:      10,
			Birthday: time.Date(2000, time.July, i%31, 0, 0, 0, 0, time.UTC),
		})
	}

	// result := db.Create(users).Statement

	//insert specific column
	result := db.Select("Name", "Age", "Birthday").Create(users).Rollback().Statement

	//insert exclude specific column
	result = db.Omit("Id", "Birthday").Create(users).Statement

	//insert in batch
	result = db.Omit("Id").CreateInBatches(users, 10).Statement

	fmt.Println(result)
}
