package comment

import (
	"time"

	"github.com/RychardProfissional/portfolio-backend/internal/core/project"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	gorm.Model
	ID        	uuid.UUID      `gorm:"column:id"`
	IP        	string         `gorm:"column:ip"`
	UserName  	string         `gorm:"column:user_name"`
	Text      	string         `gorm:"column:text"`
	ProjectID	*uuid.UUID     `gorm:"column:project_id"`
	Project     project.Entity `gorm:"column:project"`
	CreatedAt   *time.Time     `gorm:"column:created_at;default:now"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at"`
}

func (r *Repository) FromModel() (*Entitie) {
	// TODO
	return nil
}