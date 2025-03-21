package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameProject = "projects"

type Project struct {
	gorm.Model
	ID      	uuid.UUID          `gorm:"column:id;primaryKey;default:uuid_generate_v4()"`
	Title       string             `gorm:"column:title"`
	Categories	[]string           `gorm:"column:categories;type:[]text"`
	Links       []string           `gorm:"column:links;type:[]text"`
	Description string             `gorm:"column:description"`
	Image       string             `gorm:"column:image"`   
	Comments    []Comment `gorm:"foreignKey:ProjectID;references:ID"`  
	InitDate    time.Time          `gorm:"column:init_date"`
	FinishDate  time.Time          `gorm:"column:finish_date"`
	CreatedAt   *time.Time         `gorm:"column:created_at;default:now()"`
	UpdatedAt   *time.Time         `gorm:"column:updated_at;default:now()"`
}

func (*Project) TableName() string {
	return TableNameProject
}