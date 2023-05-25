package entity
import ("time")


type Commission struct {
	ID              uint64    `json:"id" gorm:"primary_key"`
	Name            string    `json:"name" binding:"required"`
	TwitterProfile  string    `json:"twitter_profile_url" binding:"required"`
	ProfilePicture  string    `json:"profile_pictModel" binding:"required"`
	LiveChatPicture string    `json:"live_chat_picture" binding:"required"`
	YoutubeUrl      string    `json:"youtube_url" binding:"required"`
	Status          string    `json:"status" binding:"required"`
	PriceListID     int       `json:"price_list_id" binding:"required"`
	PriceList       PriceList `gorm:"ForeignKey:ID;references:PriceListID"`
	ProjectID       int       `json:"project_id" binding:"required"`
	Project         Project   `gorm:"ForeignKey:ID;references:ProjectID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}