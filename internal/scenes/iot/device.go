package iot

import "github.com/google/uuid"

type Device struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	AutoUpdate bool      `json:"auto_update"`
}
