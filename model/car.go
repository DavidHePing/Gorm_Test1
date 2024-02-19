package model

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Price     float64
	UpdatedAt time.Time
}

func (car *Car) BeforeUpdate(db *gorm.DB) {
	car.UpdatedAt = time.Now()
}
