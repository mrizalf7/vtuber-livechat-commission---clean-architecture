package repository

import (
	"main/entity"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User interface {
	InsertUser(p entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
	GetRegisteredUsers() []entity.User
	// DeleteUser(p entity.User)
	// FindUserById(id uint64) entity.User
	// UpdateUser(p entity.User) entity.User
}

type UserRepository struct{
	connection *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) User{
	return &UserRepository{
		connection : dbConn,
	}
}

func (db *UserRepository) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db * UserRepository) GetRegisteredUsers() []entity.User{
	var users []entity.User
	db.connection.Order("id asc").Find(&users)
	return users
	// var p []entity.Project
	// db.connection.Order("id asc").Find(&p)
	// return p
}

func (db *UserRepository) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *UserRepository) InsertUser(p entity.User) entity.User{

	// db.connection.Save(&p)
	// db.connection.Find(&p)
	// return p
	p.Password = hashAndSalt([]byte(p.Password))
	db.connection.Save(&p)
	return p
}


// func (db *UserRepository) GetUser() []entity.User {
// 	var p []entity.User
// 	db.connection.Find(&p)
// 	return p
// }

func (db *UserRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}


func (db *UserRepository) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *UserRepository) ProfileUser(userID string) entity.User {
	var user entity.User
	// db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)
	db.connection.Find(&user,userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}


// func (db *UserRepository) UpdateUser(p entity.User) entity.User {
// 	db.connection.Save(&p)
// 	db.connection.Find(&p)
// 	return p
// }

// func (db *UserRepository) FindUserById(id uint64) entity.User {
// 	var p entity.User
// 	db.connection.Find(&p, id)
// 	return p
// }

// func (db *UserRepository) DeleteUser(p entity.User) {
// 	db.connection.Delete(&p)
// }
