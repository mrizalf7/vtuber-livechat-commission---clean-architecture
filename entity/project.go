package entity

import ("time")
// import("gorm.io/gorm")
type Project struct {
	// gorm.Model
	ID          uint64 `json:"id" gorm:"primary_key"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}