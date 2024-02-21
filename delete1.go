package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func delete_test1(db *gorm.DB) {
	car := model.Car{Id: 3}

	db.Clauses(clause.Returning{}).Delete(&car)
	fmt.Println(car)
}
