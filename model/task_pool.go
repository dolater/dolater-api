package model

import (
	"github.com/google/uuid"
)

type TaskPool struct {
	Id     uuid.UUID
	UserId string
	User   User
	Type   string // active, archived, pending
}
