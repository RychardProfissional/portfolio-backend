package comment

import (
	"errors"

	"github.com/google/uuid"
)

type Usecase struct{
	DB DB
}

func (*Usecase) ValidateCreate(comment *Entitie) (*Entitie, error) {
	comment.ID = uuid.Nil

	if comment.IP == "" {
		return nil, errors.New("ip não pode estar em falta")
	}

	if comment.ProjectID == nil {
		return nil, errors.New("project_id não pode estar em falta")
	}

	// TODO: verficar se IP existe no mesmo projectID getCommentByProjIDAndIP

	if comment.UserName == "" {
		return nil, errors.New("user_name não pode estar em falta")
	}

	if comment.Text == "" {
		return nil, errors.New("text não pode estar em falta")
	}

	comment.ID = uuid.Nil
	comment.CreatedAt = nil
	comment.UpdatedAt = nil

	return comment, nil
}

func (u *Usecase) Create(comment *Entitie) (*Entitie, error) {
	return u.DB.Create(comment)	
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
