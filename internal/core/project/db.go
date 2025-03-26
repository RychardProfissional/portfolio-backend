package project

import (
	"errors"
	"log"

	"github.com/RychardProfissional/portfolio-backend/entities"
	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"
	"github.com/RychardProfissional/portfolio-backend/models"
	"github.com/RychardProfissional/portfolio-backend/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DB struct{}

func (r *DB) Create(project *entities.Project) (*entities.Project, error) {
	model := util.ProjectToModel(project)
	res := db.Conn.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return util.ProjectFromModel(model), nil
}

func (r *DB) GetByID(id string) (*entities.Project, error) {
	var model models.Project
	res := db.Conn.Where("id = ?", id).First(&model)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return util.ProjectFromModel(&model), nil
}

func (r *DB) Update(project *entities.Project) (*entities.Project, error) {
	if project.ID == uuid.Nil {
		return nil, errors.New("id não passado")
	}
	var model models.Project
	res := db.Conn.Where("id = ?", project.ID.String()).First(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	model.Description = project.Description
	model.Categories = project.Categories
	model.Image = project.Image
	model.Title = project.Title
	model.FinishDate = project.FinishDate
	model.InitDate = project.InitDate

	log.Print(project.Description)
	log.Print(model.Description)

	res = db.Conn.Save(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return util.ProjectFromModel(&model), nil
}

func (r *DB) Delete(id string) error {
	return db.Conn.Unscoped().Where("id = ?", id).Delete(models.Project{}).Error;
}