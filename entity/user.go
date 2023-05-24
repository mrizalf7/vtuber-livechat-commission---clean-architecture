package entity


type User struct {
	// ID       uint64    `json:"id" gorm:"primary_key"`
	// Username string `json:"username" binding:"required" gorm:"unique;notnull"`
	// // Password string `json:"password" binding:"required" gorm:"notnull"`
	// Password string  `gorm:"->;<-;not null" json:"-"`
	// Token 	 string `json:"token,omitempty" gorm:"-"`
	// // Token    string  `gorm:"-" json:"token,omitempty"`
	// Email    string `json:"email" binding:"required" gorm:"unique;notnull"`
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Username     string  `gorm:"type:varchar(255)" json:"username"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`
}