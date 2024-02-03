package main

type Car struct {
	Id   int `gorm:"primaryKey"`
	Name string
}
