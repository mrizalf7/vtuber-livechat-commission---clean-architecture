package dto


type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Username     string `json:"username" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}