package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connectdb() {
	database, err := gorm.Open(mysql.Open("root:12345678@tcp(localhost:3306)/pustaka_api2"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Produk{})

	DB = database
}
