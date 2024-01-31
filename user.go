package main

import "time"

type User struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Age      int
	Birthday time.Time
}

func (User) TableName() string {
	return "test_users"
}
