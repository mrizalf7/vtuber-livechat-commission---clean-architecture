package main

import (
	"fmt"
	"main/config"
	"main/controllere"
	"main/repository"
	_"main/routes"
	"main/service"
	"github.com/gin-gonic/gin"
	_"gorm.io/gorm"
)

func main() {
	db := config.Connect()
	jwtService := service.NewJWTService()
	projectRepository := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepository)
	projectController := controllere.NewProjectController(projectService,jwtService)

	priceListRepository := repository.NewPriceListRepository(db)
	priceListService := service.NewPriceListService(priceListRepository)
	priceListController := controllere.NewPriceListController(priceListService,jwtService)

	commissionRepository := repository.NewCommissionRepository(db)
	commissionService := service.NewCommissionService(commissionRepository)
	commissionController := controllere.NewCommissionController(commissionService,jwtService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)
	
	userController := controllere.NewUserController(userService,jwtService)
	authController := controllere.NewAuthController(authService,jwtService)

	router := gin.New()
	router.GET("/",func(c *gin.Context) {
	c.String(200, "Welcome")
})
	projectController.ProjectRoutes(router)
	priceListController.PriceListRoutes(router)
	commissionController.CommissionRoutes(router)
	userController.UserRoutes(router)
	authController.AuthRoutes(router)
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// router.Use(gin.Recovery())
	// router.Use(AuthRequired())

	fmt.Println("Running on 4000")
	router.Run(":4000")
}