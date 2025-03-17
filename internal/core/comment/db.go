package comment

import "github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"

type DB struct{}

func (*DB) Create(comment *Entitie) (*Entitie, error) {
	model := comment.ToModel()
	res := db.Conn.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}

	return model.FromModel(), nil
}