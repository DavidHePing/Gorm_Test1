package main

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

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

func insert_map(db *gorm.DB) {
	result := db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	}).Statement

	fmt.Println(result.Statement)
}
