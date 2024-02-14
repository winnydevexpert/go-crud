package model

import (
	uuid "github.com/google/uuid"
)

type AppData struct {
	Id    uuid.UUID `json:"appDataId"`
	Title *string   `json:"title"`

	CategoryId uuid.UUID
}
