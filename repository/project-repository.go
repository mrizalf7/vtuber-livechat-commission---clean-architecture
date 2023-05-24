package repository

import (
	"main/entity"
	"gorm.io/gorm"
	"time"
)

type Project interface {

	InsertProject(p entity.Project) entity.Project
	GetProject() []entity.Project
	DeleteProject(p entity.Project)
	FindProjectById(id uint64) entity.Project
	UpdateProject(p entity.Project) entity.Project
}

type ProjectConnection struct{
	connection *gorm.DB
}

func NewProjectRepository(dbConn *gorm.DB) Project {
	return &ProjectConnection{
		connection: dbConn,
	}
}

func (db *ProjectConnection) InsertProject(p entity.Project) entity.Project{
	// db.connection.Create(&p)
	db.connection.Save(&p)
	db.connection.Find(&p)
	return p
}


func (db *ProjectConnection) GetProject() []entity.Project {
	var p []entity.Project
	db.connection.Order("id asc").Find(&p)
	return p
}

func (db *ProjectConnection) UpdateProject(p entity.Project) entity.Project {
	// db.connection.Updates(&p)
	// UPDATE Customers
	// SET ContactName = 'Alfred Schmidt', City= 'Frankfurt'
	// WHERE CustomerID = 1;
	db.connection.Raw("UPDATE projects SET name = ?, description = ?, updated_at = ? WHERE id = ? Returning created_at, updated_at",
	p.Name, p.Description,time.Now(),p.ID).Scan(&p)
	// db.connection.Find(&p)
	// db.connection.Take(&p)
	return p
}

func (db *ProjectConnection) FindProjectById(id uint64) entity.Project {
	var p entity.Project
	db.connection.Find(&p, id)
	return p
}

func (db *ProjectConnection) DeleteProject(p entity.Project) {
	db.connection.Delete(&p)
}