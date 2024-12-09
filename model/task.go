package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id         uuid.UUID
	UserId     string
	Title      *string
	URL        *string
	CreatedAt  time.Time
	ArchivedAt *time.Time
	DeletedAt  *time.Time
	PoolId     *uuid.UUID
	Pool       *TaskPool
}
