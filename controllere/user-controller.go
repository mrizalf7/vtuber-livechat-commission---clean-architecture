package controllere

import (
	"fmt"
	"net/http"
	"strconv"

	// "fmt"
	"main/dto"
	_"main/entity"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	// "main/entity"
	"main/helper"
	"main/service"
)

//UserController is a ...
type UserController interface {
	// All(context *gin.Context)
	// FindByID(context *gin.Context)
	// Insert(context *gin.Context)
	// Update(context *gin.Context)
	// Delete(context *gin.Context) 
	// InsertUser(context *gin.Context)
	// UpdateUser(context *gin.Context) 
	// GetUser(context *gin.Context)
	// DeleteUser(context *gin.Context)
	// FindUserById(context *gin.Context)
	UserRoutes(router *gin.Engine)
	Update(context *gin.Context)
	Profile(context *gin.Context)
	GetRegisteredUsers(context * gin.Context)
}

type userController struct {
	UserService service.UserService
	jwtService  service.JWTService
}

//NewUserController create a new instances of UserController
func NewUserController(UserServ service.UserService,jwtServ service.JWTService) UserController {
	return &userController{
		UserService : UserServ,
		jwtService:  jwtServ,
	}
}

func (c *userController) Profile(context *gin.Context){
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.UserService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}


func (c *userController) Update(context* gin.Context){
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		response := helper.BuildErrorResponse("Failed to proces request","No token found",nil)
		context.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.UserService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c* userController) GetRegisteredUsers(context *gin.Context){
	// var users [] entity
	users := c.UserService.GetRegisteredUsers()
	res:=helper.BuildResponse(true,"Sucess Get Users",users)
	context.JSON(200,res)
	return
}
//* -> note, () -> variable
//func (param) (*return function (function name)(param))

// func (c *userController) GetUser(context *gin.Context) {
// 	var User []entity.User = c.UserService.GetUser()
// 	res := helper.BuildResponse(true, "OK", User)
// 	context.JSON(http.StatusOK, res)
// }

// func (c *userController) FindUserById(context *gin.Context) {
// 	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
// 	if err != nil {
// 		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
// 		context.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	var User entity.User = c.UserService.FindUserById(id)
// 	if (User == entity.User{}) {
// 		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
// 		context.JSON(http.StatusNotFound, res)
// 	} else {
// 		res := helper.BuildResponse(true, "OK", User)
// 		context.JSON(http.StatusOK, res)
// 	}
// }

// func (c *userController) InsertUser(context *gin.Context) {
// 	// var UserCreateDTO dto.UserCreateDTO
// 	var UserCreateDTO entity.User
// 	errDTO := context.ShouldBind(&UserCreateDTO)
// 	if errDTO != nil {
// 		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
// 		context.JSON(http.StatusBadRequest, res)
// 	} else {
// 		// authHeader := context.GetHeader("Authorization")
// 		// userID := c.getUserIDByToken(authHeader)
// 		// convertedUserID, err := strconv.ParseUint(userID, 10, 64)
// 		// if err == nil {
// 		// 	UserCreateDTO.ID = convertedUserID										
// 		// }
// 		result := c.UserService.InsertUser(UserCreateDTO)
// 		response := helper.BuildResponse(true, "OK", result)
// 		context.JSON(http.StatusCreated, response)
// 	}
// }

// func (c *userController) UpdateUser(context *gin.Context) {
// 	// var UserUpdateDTO dto.UserUpdateDTO
// 	var UserUpdateDTO entity.User
// 	errDTO := context.ShouldBind(&UserUpdateDTO)
// 	if errDTO != nil {
// 		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
// 		context.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	// authHeader := context.GetHeader("Authorization")
// 	// token, errToken := c.jwtService.ValidateToken(authHeader)
// 	// if errToken != nil {
// 	// 	panic(errToken.Error())
// 	// }
// 	// claims := token.Claims.(jwt.MapClaims)
// 	// userID := fmt.Sprintf("%v", claims["user_id"])
// 	// if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
// 	// 	id, errID := strconv.ParseUint(userID, 10, 64)
// 	// 	if errID == nil {
// 	// 		bookUpdateDTO.UserID = id
// 	// 	}
// 		result := c.UserService.UpdateUser(UserUpdateDTO)
// 		response := helper.BuildResponse(true, "OK", result)
// 		context.JSON(http.StatusOK, response)
// 	// } 
// 	// else {
// 	// 	response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
// 	// 	context.JSON(http.StatusForbidden, response)
// 	// }
// }


// func (c *userController) DeleteUser(context *gin.Context) {
// 	var User entity.User
// 	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
// 	if err != nil {
// 		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
// 		context.JSON(http.StatusBadRequest, response)
// 	}
// 	User.ID = id
// 	// authHeader := context.GetHeader("Authorization")
// 	// token, errToken := c.jwtService.ValidateToken(authHeader)
// 	// if errToken != nil {
// 	// 	panic(errToken.Error())
// 	// }
// 	// claims := token.Claims.(jwt.MapClaims)
// 	// userID := fmt.Sprintf("%v", claims["user_id"])
// 	// if c.bookService.IsAllowedToEdit(userID, book.ID) {
// 		c.UserService.DeleteUser(User)
// 		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
// 		context.JSON(http.StatusOK, res)
// 	// }
// 	//  else {
// 	// 	response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
// 	// 	context.JSON(http.StatusForbidden, response)
// 	// }
// 	}

// func (c *userController) getUserIDByToken(token string) string {
// 	aToken, err := c.jwtService.ValidateToken(token)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	claims := aToken.Claims.(jwt.MapClaims)
// 	id := fmt.Sprintf("%v", claims["user_id"])
// 	return id
// }

func (c *userController) UserRoutes(router *gin.Engine){
	userRoutes := router.Group("/api/user")
	{
	// userRoutes.GET("/", c.GetUser)
	// userRoutes.POST("/", c.InsertUser)
	userRoutes.GET("/profile", c.Profile)
	userRoutes.PUT("/profile", c.Update)
	userRoutes.GET("/registered",c.GetRegisteredUsers)
	// userRoutes.DELETE("/:id" ,c.DeleteUser)
	// userRoutes.PUT("/:id", c.UpdateUser)
	}
}