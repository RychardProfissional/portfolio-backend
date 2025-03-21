package util

import (
	"github.com/RychardProfissional/portfolio-backend/entities"
	"github.com/RychardProfissional/portfolio-backend/models"
)

func CommentFromModel(m *models.Comment) (*entities.Comment) {
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
	return &models.Comment{
		ID: e.ID,
		IP: e.IP,
		UserName: e.UserName,
		Text: e.Text,
		ProjectID: e.ProjectID,
	}
}