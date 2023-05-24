package dto

//RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	// Username     string `json:"username" form:"name" binding:"required"`
	// Email    string `json:"email" form:"email" binding:"required,email" `
	// Password string `json:"password" form:"password" binding:"required"`
	Username     string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Password string `json:"password" form:"password" binding:"required"`
}
