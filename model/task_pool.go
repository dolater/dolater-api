package model

import (
	"github.com/google/uuid"
)

type TaskPool struct {
	Id    uuid.UUID
	Owner *User
	Type  string // active, archived, pending
}
