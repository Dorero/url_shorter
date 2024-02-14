package model

import "github.com/google/uuid"

type Url struct {
	Id   uuid.UUID `json:"id"`
	Path string    `json:"path"`
}
