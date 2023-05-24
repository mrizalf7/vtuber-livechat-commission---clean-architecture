package controllere

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"main/entity"
	"main/helper"
	"main/service"
	"main/middleware"
)

type ProjectController interface {
	InsertProject(context *gin.Context)
	UpdateProject(context *gin.Context) 
	GetProject(context *gin.Context) 
	DeleteProject(context *gin.Context)
	FindProjectById(context *gin.Context)
	ProjectRoutes(router *gin.Engine)
}

type projectController struct {
	ProjectService service.ProjectService
	jwtService  service.JWTService
}

//NewProjectController create a new instances of ProjectController
func NewProjectController(ProjectServ service.ProjectService,jwtServ service.JWTService) ProjectController {
	return &projectController{
		ProjectService : ProjectServ,
		jwtService:  jwtServ,
	}
}

func (c *projectController) GetProject(context *gin.Context) {
	var Project []entity.Project = c.ProjectService.GetProject()
	res := helper.BuildResponse(true, "OK", Project)
	context.JSON(http.StatusOK, res)
}

func (c *projectController) FindProjectById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var Project entity.Project = c.ProjectService.FindProjectById(id)
	var projectOnlyStruct entity.Project
	if (Project == projectOnlyStruct) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", Project)
		context.JSON(http.StatusOK, res)
	}
}

func (c *projectController) InsertProject(context *gin.Context) {
	var ProjectCreateDTO entity.Project
	errDTO := context.ShouldBind(&ProjectCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	} 
		result := c.ProjectService.InsertProject(ProjectCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

func (c *projectController) UpdateProject(context *gin.Context) {
	var ProjectUpdateDTO entity.Project
	id := context.Param("id")
	idUint64,err := strconv.ParseUint(id,10,64)
	if err!= nil {
		response := helper.BuildErrorResponse("No param id were found", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
		return
	}
	var Project entity.Project = c.ProjectService.FindProjectById(idUint64)
	var projectOnlyStruct entity.Project
	if (Project == projectOnlyStruct) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
		return
	}
	ProjectUpdateDTO.ID = idUint64
	errDTO := context.ShouldBind(&ProjectUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
		result := c.ProjectService.UpdateProject(ProjectUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
}

func (c *projectController) DeleteProject(context *gin.Context) {
	var Project entity.Project //-> EMPTY DATA ENTITY, ONLY The STRUCT/schema/table
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param id were found", err.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
		return
	}
	selectedProject := c.ProjectService.FindProjectById(id)
	if(selectedProject == Project){
		res:= helper.BuildErrorResponse("Data not found","No data with given id",helper.EmptyObj{})
		context.JSON(http.StatusBadRequest,res)
		return
	}
	Project.ID = id
	c.ProjectService.DeleteProject(Project)
	res := helper.BuildResponse(true, "Deleted", selectedProject)
	context.JSON(http.StatusOK, res)
}

func (c *projectController) ProjectRoutes(router *gin.Engine){
	// jwtService := service.NewJWTService()
	// projectRoutes := router.Group("/api/project",middleware.AuthorizeJWT(jwtService))
	projectRoutes := router.Group("/api/project",middleware.AuthorizeJWT(c.jwtService))
	{
	projectRoutes.POST("/", c.InsertProject)
	projectRoutes.DELETE("/:id" ,c.DeleteProject)
	projectRoutes.PUT("/:id", c.UpdateProject)
	}
	router.GET("/api/project", c.GetProject)
	router.GET("/api/project/:id", c.FindProjectById)
}