package model

import (
	uuid "github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	MidName   *string   `json:"midName"`
	LastName  string    `json:"lastName"`
}
