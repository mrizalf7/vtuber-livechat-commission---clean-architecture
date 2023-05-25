package repository

import (
	// "github.com/ydhnwb/golang_api/entity"
	"main/entity"

	"gorm.io/gorm"
)

type PriceList interface {
	InsertPriceList(p entity.PriceList) entity.PriceList
	GetPriceList() []entity.PriceList
	DeletePriceList(p entity.PriceList)
	FindPriceListById(id uint64) entity.PriceList
	UpdatePriceList(p entity.PriceList) entity.PriceList
	DeletePriceListByCategory(p entity.PriceList)
}

type priceListConnection struct {
	connection *gorm.DB
}

func NewPriceListRepository(dbConn *gorm.DB) PriceList {
	return &priceListConnection{
		connection: dbConn,
	}
}

func (db *priceListConnection) InsertPriceList(p entity.PriceList) entity.PriceList {

	// db.connection.Save(&p)
	db.connection.Create(&p)
	db.connection.Preload("Project").Find(&p)
	return p
}

func (db *priceListConnection) GetPriceList() []entity.PriceList {
	var p []entity.PriceList
	db.connection.Order("id asc").Preload("Project").Find(&p)
	return p
}

func (db *priceListConnection) UpdatePriceList(p entity.PriceList) entity.PriceList {
	// db.connection.Save(p)
	db.connection.Updates(&p)
	db.connection.Preload("Project").Find(&p)
	return p
}

func (db *priceListConnection) FindPriceListById(id uint64) entity.PriceList {
	var p entity.PriceList
	db.connection.Preload("Project").Find(&p, id)
	return p
}

func (db *priceListConnection) DeletePriceList(p entity.PriceList) {
	db.connection.Delete(&p)
}

func (db *priceListConnection) DeletePriceListByCategory(p entity.PriceList) {
	// db.connection.Delete(&p)
	// db.connection.Where("category = ?", p.Category).Delete(&p)
	db.connection.Raw("Delete from price_lists where category=?",p.Category).Scan(&p)

}
