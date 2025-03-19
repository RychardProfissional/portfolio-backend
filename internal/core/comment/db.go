package comment

import (
	"errors"

	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"
	"gorm.io/gorm"
)

type DB struct{}

func (*DB) Create(comment *Entitie) (*Entitie, error) {
	model := comment.ToModel()
	res := db.Conn.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return model.FromModel(), nil
}

func (*DB) GetByID(id string) (*Entitie, error){
	var model *Repository
	res := db.Conn.Where("id = ?", id).First(&model)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return model.FromModel(), nil
}

func (*DB) UpdateText(id string, text *string) (error) {
	if id == "" {
		return errors.New("id not found")
	}
	if text == nil {
		return errors.New("text not found")
	}

	res := db.Conn.
		Model(&Repository{}).
		Where("id = ?", id).
		Update("text", *text)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (*DB) Delete(id string) error {
	return db.Conn.Unscoped().Where("id = ?", id).Delete(Repository{}).Error;
}