package controllere

import (
	// "fmt"
	"net/http"
	"strconv"
	"fmt"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/dto"
	"main/entity"
	"main/helper"
	"main/service"
	"main/middleware"
)

//PriceListController is a ...
type PriceListController interface {
	InsertPriceList(context *gin.Context)
	UpdatePriceList(context *gin.Context) 
	GetPriceList(context *gin.Context) 
	DeletePriceList(context *gin.Context)
	FindPriceListById(context *gin.Context)
	DeletePriceListByCategory(context *gin.Context)
	PriceListRoutes(router *gin.Engine)
}

type priceListController struct {
	priceListService service.PriceListService
	jwtService  service.JWTService
}

//NewPriceListController create a new instances of PriceListController
func NewPriceListController(priceListServ service.PriceListService,jwtServ service.JWTService) PriceListController {
	return &priceListController{
		priceListService : priceListServ,
		jwtService:  jwtServ,
	}
}
//* -> note, () -> variable
//func (param) (*return function (function name)(param))

func (c *priceListController) GetPriceList(context *gin.Context) {
	var priceList []entity.PriceList = c.priceListService.GetPriceList()
	res := helper.BuildResponse(true, "OK", priceList)
	context.JSON(http.StatusOK, res)
}

func (c *priceListController) FindPriceListById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var priceList entity.PriceList = c.priceListService.FindPriceListById(id)
	if (priceList == entity.PriceList{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", priceList)
		context.JSON(http.StatusOK, res)
	}
}

func (c *priceListController) InsertPriceList(context *gin.Context) {
	var priceListCreateDTO dto.PriceListCreateDTO
	errDTO := context.ShouldBind(&priceListCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.priceListService.InsertPriceList(priceListCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *priceListController) UpdatePriceList(context *gin.Context) {
	var priceListUpdateDTO dto.PriceListUpdateDTO

	id := context.Param("id")
	idUint64,err := strconv.ParseUint(id,10,64)
	if err!= nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	priceListUpdateDTO.ID = idUint64

	errDTO := context.ShouldBind(&priceListUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

		result := c.priceListService.UpdatePriceList(priceListUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
}

func (c *priceListController) DeletePriceListByCategory(context *gin.Context){
	var priceList entity.PriceList
	category := context.Param("category")
	fmt.Println("Category name:"+category)
	// if err != nil {
	// 	response := helper.BuildErrorResponse("Failed to get name parameter", " no param name were found",helper.EmptyObj{})
	// 	context.JSON(http.StatusBadRequest,response)
	// }
	priceList.Category = category
	c.priceListService.DeletePriceListByCategory(priceList)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK,res)
}

func (c *priceListController) DeletePriceList(context *gin.Context) {
	var priceList entity.PriceList
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
		return
	}
	selectedPriceList := c.priceListService.FindPriceListById(id)
	if (selectedPriceList == priceList){
		res:=helper.BuildErrorResponse("Data not found","No data with given id",helper.EmptyObj{})
		context.JSON(http.StatusNotFound,res)
		return
	}
	priceList.ID = id
	c.priceListService.DeletePriceList(priceList)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}

func (c *priceListController) PriceListRoutes(router *gin.Engine){
	priceListRoutes := router.Group("api/pricelist",middleware.AuthorizeJWT(c.jwtService))
	{
		priceListRoutes.POST("/", c.InsertPriceList)
		priceListRoutes.DELETE("/:id", c.DeletePriceList)
		priceListRoutes.PUT("/:id", c.UpdatePriceList)
		// priceListRoutes.DELETE("/name/:category", c.DeletePriceListByCategory)
	}
	router.GET("api/pricelist", c.GetPriceList)
	router.GET("api/pricelist/:id", c.FindPriceListById)
}