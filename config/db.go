package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/entity"
)


func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://rizal:rizalroot@localhost:5432/postgres"), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	db.AutoMigrate(&entity.PriceList{})
	db.AutoMigrate(&entity.Commission{})
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Project{})
	return db
}