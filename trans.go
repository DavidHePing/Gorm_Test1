package main

import (
	"Gorm_Test1/model"
	"fmt"

	"gorm.io/gorm"
)

func tran_test1(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		//would rollback if any error
		if err := tx.Model(&model.Car{Id: 2}).Update("price = ?", 200).Error; err != nil {
			return err
		}

		if err := tx.Model(model.Car{}).Where("name = ?", 1).Update("price = ?", 200).Error; err != nil {
			return err
		}

		return nil
	})

	car := model.Car{Id: 2}
	db.Find(&car)
	fmt.Println(car.Price)
}
