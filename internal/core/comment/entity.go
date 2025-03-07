package comment

import (
	"github.com/RychardProfissional/portfolio-backend/internal/core/project"
	"github.com/google/uuid"
)

type Entitie struct {
	ID        	uuid.UUID      `json:"id"`
	IP        	string         `json:"ip"`
	UserName  	string         `json:"user_name"`
	Text      	string         `json:"text"`
	ProjectID	*uuid.UUID     `json:"project_id"`
	Project     project.Entity `json:"project,omitempty"`
}

func (e *Entitie) ToModel() *Repository {
	// TODO
	return nil
}