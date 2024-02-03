package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

	select_test1(db)
	// insert_test1(db)
	// insert_map(db)
}
