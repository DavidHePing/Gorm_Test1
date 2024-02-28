package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func update_where_Test1(db *gorm.DB) {
	carToyota := model.Car{Price: 155}
	db.Model(&carToyota).Clauses(clause.Returning{}).Where("name = ?", "Toyota").Select("price").Updates(carToyota)

	fmt.Println(carToyota)
}

func update_where_Test2(db *gorm.DB) {
	carToyota := model.Car{Price: 155}
	result := db.Model(&carToyota).Clauses(clause.Returning{}).Where("name = ?", 1).Select("price").Updates(carToyota)

	fmt.Println(result.Error)
}
