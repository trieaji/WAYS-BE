package database

import (
	"fmt"
	"ways/models"
	"ways/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Product{},
		&models.Cart{},
		&models.CartTopping{},
		&models.Topping{},
		&models.Transaction{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
