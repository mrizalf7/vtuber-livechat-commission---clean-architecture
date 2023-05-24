package service

import(
	_"fmt"
	"log"
	"github.com/mashingan/smapping"
	"main/entity"
	"main/repository"
)


type ProjectService interface {
	InsertProject(p entity.Project) entity.Project
	UpdateProject(p entity.Project) entity.Project
	GetProject() []entity.Project
	DeleteProject(p entity.Project)
	FindProjectById(id uint64) entity.Project
	// IsAllowedToEdit(userID string, bookID uint64) bool
}

type projectService struct {
	ProjectRepository repository.Project
}

//NewProjectService .....
func NewProjectService(ProjectRepo repository.Project) ProjectService {
	return &projectService{
		ProjectRepository: ProjectRepo,
	}
}
//Project
func (service *projectService) InsertProject(p entity.Project) entity.Project {
	Project := entity.Project{}
	err := smapping.FillStruct(&Project, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.ProjectRepository.InsertProject(Project)
	return res
}

func (service *projectService) UpdateProject(p entity.Project) entity.Project {
	Project := entity.Project{}
	err := smapping.FillStruct(&Project, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.ProjectRepository.UpdateProject(Project)
	return res
}

func (service *projectService) DeleteProject(p entity.Project) {
	service.ProjectRepository.DeleteProject(p)
}

// func (service *ProjectService) DeleteProjectByCategory(p entity.Project,category_name string){
// 	service.ProjectRepository.DeleteProjectByCategory(p,category_name)
// }

func (service *projectService) GetProject() []entity.Project {
	return service.ProjectRepository.GetProject()
}

func (service *projectService) FindProjectById(id uint64) entity.Project {
	// fmt.Sprintf("%v", id) 
	return service.ProjectRepository.FindProjectById(id)
}