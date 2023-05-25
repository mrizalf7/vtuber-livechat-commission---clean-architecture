package dto


type PriceListUpdateDTO struct {

	ID          uint64  `json:"id" gorm:"primary_key"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"required"`
	PriceIDR    string  `json:"price_idr" binding:"required"`
	PriceUSD    string  `json:"price_usd" binding:"required"`
	ProjectID   int     `json:"project_id" binding:"required"`

		// gorm.Model
		// ID          int     `json:"price_list_id" gorm:"primary_key"`
		// Category    string  `json:"category"`
		// Description string  `json:"description"`
		// PriceIDR    string  `json:"price_idr"`
		// PriceUSD    string  `json:"price_usd"`
		// ProjectID   int     `json:"project_id"`
		// Project     Project `gorm:"ForeignKey:ID;references:ProjectID"`
}

type PriceListCreateDTO struct {
	// Title       string `json:"title" form:"title" binding:"required"`
	// Description string `json:"description" form:"description" binding:"required"`
	// UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"required"`
	PriceIDR    string  `json:"price_idr" binding:"required"`
	PriceUSD    string  `json:"price_usd" binding:"required"`
	ProjectID   int     `json:"project_id" binding:"required"`

		// gorm.Model
		// ID          int     `json:"price_list_id" gorm:"primary_key"`
		// Category    string  `json:"category"`
		// Description string  `json:"description"`
		// PriceIDR    string  `json:"price_idr"`
		// PriceUSD    string  `json:"price_usd"`
		// ProjectID   int     `json:"project_id"`
		// Project     Project `gorm:"ForeignKey:ID;references:ProjectID"`
}
