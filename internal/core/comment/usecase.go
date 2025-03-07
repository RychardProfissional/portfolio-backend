package comment

import (
	"errors"

	"github.com/RychardProfissional/portfolio-backend/internal/core/adapters/db"
	"github.com/google/uuid"
)

type Usecase struct{}

func (*Usecase) ValidateCreate(comment *Entitie) (*Entitie, error) {
	comment.ID = uuid.Nil

	if comment.IP == "" {
		return nil, errors.New("ip não pode estar em falta")
	}

	if *comment.ProjectID == uuid.Nil {
		return nil, errors.New("project_id não pode estar em falta")
	}

	// TODO: verficar se IP existe no mesmo projectID getCommentByProjIDAndIP

	if comment.UserName == "" {
		return nil, errors.New("user_name não pode estar em falta")
	}

	if comment.Text == "" {
		return nil, errors.New("text não pode estar em falta")
	}

	return comment, nil
}

func (*Usecase) Create(comment *Entitie) (*Entitie, error) {
	dbconn := db.Connect()

	model := comment.ToModel()
	res := dbconn.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	
	return model.FromModel(), nil
}

func (*Usecase) GetByID() {
	// TODO
}

func (*Usecase) GetByProjectID() {
	// TODO
}

func (*Usecase) ValidateUpdateText() {
	// TODO
}

func (*Usecase) UpdateText() {
	// TODO
}

func (*Usecase) Delete() {
	// TODO
}
