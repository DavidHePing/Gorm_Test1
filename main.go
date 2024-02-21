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

	// insert_test1(db)
	// insert_map(db)

	select_test1(db)
	// select_with_id(db)
	// select_where_test1(db)
	// select_where_and_test1(db)
	// select_count_test1(db)
	// select_raw_sql_test1(db)
	// select_limit_offset_test1(db)
	// select_preload_test1(db)
	// select_preload_test2(db)
	// save_test1(db)
	// udpate_test1(db)
	// udpate_test2(db)
	delete_test1(db)
}
