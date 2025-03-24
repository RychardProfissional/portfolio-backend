package util

import (
	"github.com/RychardProfissional/portfolio-backend/entities"
	"github.com/RychardProfissional/portfolio-backend/models"
)

func CommentFromModel(m *models.Comment) (*entities.Comment) {
	if m == nil {
		return nil
	}
	return &entities.Comment{
		ID: m.ID,
		IP: m.IP,
		UserName: m.UserName,
		Text: m.Text,
		ProjectID: m.ProjectID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func CommentToModel(e *entities.Comment) (*models.Comment) {
	if e == nil {
		return nil
	}
	return &models.Comment{
		ID: e.ID,
		IP: e.IP,
		UserName: e.UserName,
		Text: e.Text,
		ProjectID: e.ProjectID,
	}
}

func ProjectToModel(e *entities.Project) (*models.Project) {
	if e == nil {
		return nil
	}
	return &models.Project{
		ID: e.ID,
		Title: e.Title,
		Categories: e.Categories,
		Links: e.Links,
		Description: e.Description,
		Image: e.Image,
		Comments: CommentsToModels(e.Comments),
		InitDate: e.InitDate,
		FinishDate: e.FinishDate,
	}
}

func ProjectFromModel(m *models.Project) (*entities.Project) {
	if m == nil {
		return nil
	}
	return &entities.Project{
		ID: m.ID,
		Title: m.Title,
		Categories: m.Categories,
		Links: m.Links,
		Description: m.Description,
		Image: m.Image,
		Comments: CommentsFromModels(m.Comments),
		InitDate: m.InitDate,
		FinishDate: m.FinishDate,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}