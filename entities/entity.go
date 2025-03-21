package entities

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        	uuid.UUID       `json:"id"`
	IP        	string          `json:"ip"`
	UserName  	string          `json:"user_name"`
	Text      	string          `json:"text"`
	ProjectID	*uuid.UUID      `json:"project_id"`
	// Project     *project.Entity `json:"project,omitempty"`
	CreatedAt   *time.Time      `json:"created_at"`
 	UpdatedAt   *time.Time      `json:"updated_at"`
}
