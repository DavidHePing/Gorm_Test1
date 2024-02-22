package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Car2 struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Price     float64
	IsDeleted bool
}

func (Car2) TableName() string {
	return "test_cars"
}

func (car Car2) Select(db *gorm.DB) {
	c := car

	db.Find(&c)
	fmt.Println(c)
}

func (Car2) SelectAll(db *gorm.DB) {
	var c = []Car2{}

	db.Where("is_deleted is null").Find(&c)
	fmt.Println(c)
}

func (car Car2) Delete(db *gorm.DB) {
	var c = car

	db.Save(&c)
}
