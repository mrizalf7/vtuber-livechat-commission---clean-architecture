package repository

import (
	// "github.com/ydhnwb/golang_api/entity"
	"main/entity"
	"gorm.io/gorm"
)

type Commission interface {
	InsertCommission(p entity.Commission) entity.Commission
	GetCommission() []entity.Commission
	DeleteCommission(p entity.Commission)
	FindCommissionById(id uint64) entity.Commission
	UpdateCommission(p entity.Commission) entity.Commission
}

type CommissionConnection struct {
	connection *gorm.DB
}

func NewCommissionRepository(dbConn *gorm.DB) Commission {
	return &CommissionConnection{
		connection: dbConn,
	}
}

func (db *CommissionConnection) InsertCommission(p entity.Commission) entity.Commission {

	// db.connection.Save(&p)
	db.connection.Create(&p)
	// db.connection.Preload("Project").Preload("PriceList.Project").Find(&p)
	db.connection.Preload("PriceList.Project").Preload("Project").Find(&p)
	// db.connection.Preload("Project").Find(&p)
	return p
}

func (db *CommissionConnection) GetCommission() []entity.Commission {
	var p []entity.Commission
	db.connection.Order("id asc").Preload("Project").Preload("PriceList.Project").Find(&p)
	return p
}

func (db *CommissionConnection) UpdateCommission(p entity.Commission) entity.Commission {
	db.connection.Updates(&p)
	db.connection.Preload("Project").Preload("PriceList.Project").Find(&p)
	return p
}

func (db *CommissionConnection) FindCommissionById(id uint64) entity.Commission {
	var p entity.Commission
	db.connection.Preload("Project").Preload("PriceList.Project").Find(&p, id)
	return p
}

func (db *CommissionConnection) DeleteCommission(p entity.Commission) {
	db.connection.Delete(&p)
}