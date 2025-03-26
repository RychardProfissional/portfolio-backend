package entities

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Categories  []string   `json:"categories"`
	Links       []string   `json:"links"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	Comments    []*Comment `json:"comments,omitempty"`
	InitDate    time.Time  `json:"init_date"`
	FinishDate  time.Time  `json:"finish_date"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}