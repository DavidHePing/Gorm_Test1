package model

type Car struct {
	Id   int `gorm:"primaryKey"`
	Name string
}
