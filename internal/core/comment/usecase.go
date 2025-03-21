package comment

import (
	"errors"

	"github.com/RychardProfissional/portfolio-backend/entities"
	"github.com/google/uuid"
)

type Usecase struct{
	DB DB
}

func (*Usecase) ValidateCreate(comment *entities.Comment) (*entities.Comment, error) {
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

func (u *Usecase) Create(comment *entities.Comment) (*entities.Comment, error) {
	return u.DB.Create(comment)	
}

func (u *Usecase) GetByID(id string) (*entities.Comment, error) {
	return u.DB.GetByID(id)
}

func (*Usecase) GetByProjectID() {
	// TODO
}

func (u *Usecase) ValidateUpdateText(id, ip, text string) (*entities.Comment, error) {
	commentUUID, err := uuid.Parse(id)
	if commentUUID == uuid.Nil || err != nil  {
		return nil, errors.New("commentID invalido")
	}

	if ip == "" {
		return nil, errors.New("ip invalido")
	}

	if text == "" {
		return nil, errors.New("texto está vazio")
	}

	// TODO: converter para internal error
	comment, err := u.DB.GetByID(id)
	if err != nil {
		return nil, err
	}

	if comment.IP != ip {
		return nil, errors.New("ip não correspondente")
	}	

	return comment, nil
}

func (u *Usecase) UpdateText(id string, text *string) (error) {
	return u.DB.UpdateText(id, text)
}

func (u *Usecase) Delete(id string) error {
	return u.DB.Delete(id)
}
