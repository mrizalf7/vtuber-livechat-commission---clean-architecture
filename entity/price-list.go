package entity

import ("time")

type PriceList struct {
	// gorm.Model
	ID          uint64   `json:"id" gorm:"primary_key"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"required"`
	PriceIDR    string  `json:"price_idr" binding:"required"`
	PriceUSD    string  `json:"price_usd" binding:"required"`
	ProjectID   int     `json:"project_id" binding:"required"`
	Project     Project `gorm:"ForeignKey:ID;references:ProjectID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}