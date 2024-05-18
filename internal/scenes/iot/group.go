package iot

import "github.com/google/uuid"

type Group struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Devices []Device  `json:"devices,omitempty"`
}
