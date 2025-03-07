package project

import "github.com/google/uuid"

type Entity struct {
	ID uuid.UUID `json:"id"`
}