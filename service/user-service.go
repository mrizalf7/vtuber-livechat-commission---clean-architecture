package service

import(

	// "fmt"
	"log"
	"github.com/mashingan/smapping"
	"main/entity"
	"main/repository"
	"main/dto"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
	GetRegisteredUsers() []entity.User
	// InsertUser(p entity.User) entity.User
	// UpdateUser(p entity.User) entity.User
	// DeleteUser(p entity.User)
	// FindUserById(id uint64) entity.User
	// DeleteUserByCategory(p entity.User,category_name string)
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type UserServ struct {
	UserRepository repository.User
}

//NewUserService .....
func NewUserService(UserRepo repository.User) UserService {
	return &UserServ{
		UserRepository: UserRepo,
	}
}

func (service *UserServ) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.UserRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *UserServ) Profile(userID string) entity.User {
	return service.UserRepository.ProfileUser(userID)
}

func (service *UserServ) GetRegisteredUsers() []entity.User{
	return service.UserRepository.GetRegisteredUsers()
}
//User
// func (service *UserServ) InsertUser(p entity.User) entity.User {
// 	User := entity.User{}
// 	err := smapping.FillStruct(&User, smapping.MapFields(&p))
// 	if err != nil {
// 		log.Fatalf("Failed map %v: ", err)
// 	}
// 	res := service.UserRepository.InsertUser(User)
// 	return res
// }

// func (service *UserServ) UpdateUser(p entity.User) entity.User {
// 	User := entity.User{}
// 	err := smapping.FillStruct(&User, smapping.MapFields(&p))
// 	if err != nil {
// 		log.Fatalf("Failed map %v: ", err)
// 	}
// 	res := service.UserRepository.UpdateUser(User)
// 	return res
// }

// func (service *UserServ) DeleteUser(p entity.User) {
// 	service.UserRepository.DeleteUser(p)
// }

// func (service *UserServ) GetUser() []entity.User {
// 	return service.UserRepository.GetUser()
// }

// func (service *UserServ) FindUserById(id uint64) entity.User {
// 	fmt.Sprintf("%v", id) 
// 	return service.UserRepository.FindUserById(id)
// }