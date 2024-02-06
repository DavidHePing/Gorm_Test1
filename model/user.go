package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           int `gorm:"primaryKey"`
	Name         string
	Age          int
	Birthday     time.Time
	RandomNumber int
	Tasks        []Task
}

func (User) TableName() string {
	return "test_users"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.RandomNumber = 123
	return
}
