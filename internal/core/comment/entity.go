package comment

import (
	"time"

	"github.com/RychardProfissional/portfolio-backend/internal/core/project"
	"github.com/google/uuid"
)

type Entitie struct {
	ID        	uuid.UUID       `json:"id"`
	IP        	string          `json:"ip"`
	UserName  	string          `json:"user_name"`
	Text      	string          `json:"text"`
	ProjectID	*uuid.UUID      `json:"project_id"`
	Project     *project.Entity `json:"project,omitempty"`
	CreatedAt   *time.Time      `json:"created_at"`
 	UpdatedAt   *time.Time      `json:"updated_at"`
}

func (e *Entitie) ToModel() *Repository {
	if e == nil {
		return nil
	}

	return &Repository{
		ID:        e.ID,
		IP:        e.IP,
		UserName:  e.UserName,
		Text:      e.Text,
		ProjectID: e.ProjectID,
		// Project:   e.Project,
	}
}
