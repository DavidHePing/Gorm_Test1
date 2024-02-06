package model

type Task struct {
	Id          int `gorm:"primaryKey"`
	UserId      int
	Subject     string
	Description string
}
