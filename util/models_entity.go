package util

import (
	"github.com/RychardProfissional/portfolio-backend/entities"
	"github.com/RychardProfissional/portfolio-backend/models"
)

func CommentsToModels(e []*entities.Comment) []*models.Comment {
	var models []*models.Comment
	for _, model := range(e) {
		models = append(models, CommentToModel(model))
	}
	return models
}

func CommentsFromModels(m []*models.Comment) []*entities.Comment {
	var models []*entities.Comment
	for _, model := range(m) {
		models = append(models, CommentFromModel(model))
	}
	return models
}
