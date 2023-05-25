package dto

type CommissionCreateDTO struct{
	Name            string    `json:"name" binding:"required"`
	TwitterProfile  string    `json:"twitter_profile_url" binding:"required"`
	ProfilePicture  string    `json:"profile_pictModel" binding:"required"`
	LiveChatPicture string    `json:"live_chat_picture" binding:"required"`
	YoutubeUrl      string    `json:"youtube_url" binding:"required"`
	Status          string    `json:"status" binding:"required"`
	ProjectID       int       `json:"project_id" binding:"required"`
	PriceListID 	int 	  `json:"price_list_id" binding:"required"`
}

type CommissionUpdateDTO struct{
	ID				uint64	  `json:"id" gorm:"primary_key"`
	Name            string    `json:"name" binding:"required"`
	TwitterProfile  string    `json:"twitter_profile_url" binding:"required"`
	ProfilePicture  string    `json:"profile_pictModel" binding:"required"`
	LiveChatPicture string    `json:"live_chat_picture" binding:"required"`
	YoutubeUrl      string    `json:"youtube_url" binding:"required"`
	Status          string    `json:"status" binding:"required"`
	ProjectID       int       `json:"project_id" binding:"required"`
	PriceListID 	int 	  `json:"price_list_id" binding:"required"`
}