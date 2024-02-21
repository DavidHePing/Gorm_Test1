package model

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Id    int `gorm:"primaryKey"`
	Name  string
	Price float64
	Type  string
}

func (car *Car) BeforeUpdate(db *gorm.DB) {
	car.UpdatedAt = time.Now()
}
