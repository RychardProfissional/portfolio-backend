package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameComment = "comments"

type Comment struct {
	gorm.Model
	ID        	uuid.UUID      `gorm:"column:id;primaryKey;default:uuid_generate_v4()"`
	IP        	string         `gorm:"column:ip"`
	UserName  	string         `gorm:"column:user_name"`
	Text      	string         `gorm:"column:text"`
	ProjectID	*uuid.UUID     `gorm:"column:project_id"`
	// Project     project.Entity `gorm:"column:project"` // TODO
	CreatedAt   *time.Time     `gorm:"column:created_at;default:now()"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at;default:now()"`
}

func (r *Comment) TableName() string {
	return TableNameComment
}
