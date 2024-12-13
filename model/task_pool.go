package model

import (
	"github.com/google/uuid"
)

type TaskPool struct {
	Id      uuid.UUID
	OwnerId string
	Owner   User
	Type    string // active, archived, pending, bin
}
