package project

import (
	"errors"

	"github.com/RychardProfissional/portfolio-backend/entities"
)

type Usecase struct{
	DB DB
}

func (u *Usecase) ValidateCreate(project *entities.Project) (*entities.Project, error) {
	if project.Title == "" {
		return nil, errors.New("titulo vazio")
	}

	if project.Description == "" {
		return nil, errors.New("descrição vazia")
	}

	if project.Image == "" {
		return nil, errors.New("imagem vazia")
	}

	if project.InitDate.IsZero() {
		return nil, errors.New("data de inicio igual a zero")
	}

	if project.FinishDate.IsZero() {
		return nil, errors.New("data final igual a zero")
	}

	return project, nil
}

func (u *Usecase) Create(project *entities.Project) (*entities.Project, error) {
	return u.DB.Create(project)
}

func (u *Usecase) GetByID(id string) (*entities.Project, error) {
	return u.DB.GetByID(id)
}

func (u *Usecase) ValidateUpdate(id string, project *entities.Project) (*entities.Project, error) {
	if project.Title == "" {
		return nil, errors.New("titulo vazio")
	}

	if project.Description == "" {
		return nil, errors.New("descrição vazia")
	}

	if project.Image == "" {
		return nil, errors.New("imagem vazia")
	}

	if project.InitDate.IsZero() {
		return nil, errors.New("data de inicio igual a zero")
	}

	if project.FinishDate.IsZero() {
		return nil, errors.New("data final igual a zero")
	}

	return u.DB.GetByID(id)
}

func (u *Usecase) Update(project *entities.Project) (*entities.Project, error) {
	return u.DB.Update(project)
}

func (u *Usecase) Delete(id string) error {
	return u.DB.Delete(id)
}
