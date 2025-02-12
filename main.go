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
		//skip transaction, performance 30% up!
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Println("Failed!!!!!")
	}

	// insert_test1(db)
	// insert_map(db)

	// select_test1(db)
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
	// update_where_Test1(db)
	// update_where_Test2(db)
	// update_where_Test3(db)

	// delete_test1(db)

	// select_many_test1(db)
	// select_many_test2(db)

	// chain_test1(db)

	tran_test1(db)

	////raw sql
	// raw_select_test1(db)

	////car2
	// test_car2(db)
}
