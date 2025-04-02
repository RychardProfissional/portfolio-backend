package comment

import (
	"errors"

	"github.com/RychardProfissional/portfolio-backend/entities"
	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"
	"github.com/RychardProfissional/portfolio-backend/models"
	"github.com/RychardProfissional/portfolio-backend/util"
	"gorm.io/gorm"
)

type DB struct{}

func (*DB) Create(comment *entities.Comment) (*entities.Comment, error) {
	model := util.CommentToModel(comment)
	res := db.Conn.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return util.CommentFromModel(model), nil
}

func (*DB) VerifiedUnique(ip string, project_id string) (bool, error) {
	var model *models.Comment
	res := db.Conn.Where("ip = ? and project_id = ?", ip, project_id).First(&model)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return true, nil
		}
		return false, res.Error
	}

	return false, nil
}

func (*DB) GetByID(id string) (*entities.Comment, error){
	var model *models.Comment
	res := db.Conn.Where("id = ?", id).First(&model)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return util.CommentFromModel(model), nil
}

func (*DB) GetByProjectID(project_id string) ([]*entities.Comment, error) {
	var model []*models.Comment
	res := db.Conn.Where("project_id = ?", project_id).Find(&model)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return util.CommentsFromModels(model), nil
}

func (*DB) UpdateText(id string, text *string) (error) {
	if id == "" {
		return errors.New("id not found")
	}
	if text == nil {
		return errors.New("text not found")
	}

	res := db.Conn.
		Model(&models.Comment{}).
		Where("id = ?", id).
		Update("text", *text)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (*DB) Delete(id string) error {
	return db.Conn.Unscoped().Where("id = ?", id).Delete(models.Comment{}).Error;
}