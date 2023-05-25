package controllere

import (
	// "fmt"
	"net/http"
	"strconv"

	// "fmt"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/dto"
	"main/entity"
	"main/helper"
	"main/middleware"
	"main/service"
)

//CommissionController is a ...
type CommissionController interface {
	InsertCommission(context *gin.Context)
	UpdateCommission(context *gin.Context) 
	GetCommission(context *gin.Context) 
	DeleteCommission(context *gin.Context)
	FindCommissionById(context *gin.Context)
	CommissionRoutes(router *gin.Engine)
}

type commissionController struct {
	CommissionService service.CommissionService
	jwtService  service.JWTService
}

//NewCommissionController create a new instances of CommissionController
func NewCommissionController(CommissionServ service.CommissionService,jwtServ service.JWTService) CommissionController {
	return &commissionController{
		CommissionService : CommissionServ,
		jwtService:  jwtServ,
	}
}
//* -> note, () -> variable
//func (param) (*return function (function name)(param))

func (c *commissionController) GetCommission(context *gin.Context) {
	var Commission []entity.Commission = c.CommissionService.GetCommission()
	res := helper.BuildResponse(true, "OK", Commission)
	context.JSON(http.StatusOK, res)
}

func (c *commissionController) FindCommissionById(context *gin.Context) {
	// var CommissionFindByIdDTO entity.Commission
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var selectedCommission entity.Commission = c.CommissionService.FindCommissionById(id)
	if (selectedCommission == entity.Commission{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", selectedCommission)
		context.JSON(http.StatusOK, res)
	}
}

func (c *commissionController) InsertCommission(context *gin.Context) {
	var CommissionCreateDTO dto.CommissionCreateDTO
	// var CommissionCreateDTO entity.Commission
	errDTO := context.ShouldBind(&CommissionCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
		result := c.CommissionService.InsertCommission(CommissionCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

func (c *commissionController) UpdateCommission(context *gin.Context) {
	var CommissionUpdateDTO dto.CommissionUpdateDTO
	// var CommissionUpdateDTO entity.Commission
	id := context.Param("id")
	idUint64,err := strconv.ParseUint(id,10,64)
	if err!=nil{
		response := helper.BuildErrorResponse("No param id were found",err.Error(),helper.EmptyObj{})
		context.JSON(http.StatusBadRequest,response)
		return
	}
	selectedCommission := c.CommissionService.FindCommissionById(idUint64)
	var commissionOnlyStruct entity.Commission
	if(selectedCommission == commissionOnlyStruct){
		res:= helper.BuildErrorResponse("Data not found","No data with given id",helper.EmptyObj{})
		context.JSON(http.StatusNotFound,res)
		return
	}
	CommissionUpdateDTO.ID = idUint64
	errDTO := context.ShouldBind(&CommissionUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
		result := c.CommissionService.UpdateCommission(CommissionUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
}

func (c *commissionController) DeleteCommission(context *gin.Context) {
	var CommissionDeleteDTO entity.Commission
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	selectedCommission := c.CommissionService.FindCommissionById(id)
	var commissionOnlyStruct entity.Commission 
	if (selectedCommission == commissionOnlyStruct){
		res:= helper.BuildErrorResponse("Data not found","No data with given id",helper.EmptyObj{})
		context.JSON(http.StatusNotFound,res)
		return
	}
	CommissionDeleteDTO.ID = id
	c.CommissionService.DeleteCommission(CommissionDeleteDTO)
	res := helper.BuildResponse(true, "Deleted", selectedCommission)
	context.JSON(http.StatusOK, res)
}

func (c *commissionController) CommissionRoutes(router *gin.Engine){
	CommissionRoute := router.Group("api/commission",middleware.AuthorizeJWT(c.jwtService))
	{
	CommissionRoute.POST("/", c.InsertCommission)
	CommissionRoute.DELETE("/:id", c.DeleteCommission)
	CommissionRoute.PUT("/:id", c.UpdateCommission)
	}
	router.GET("api/commission", c.GetCommission)
	router.GET("api/commission/:id", c.FindCommissionById)
}