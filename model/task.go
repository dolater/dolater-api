package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	URL         *string
	CreatedAt   time.Time
	CompletedAt *time.Time
	ArchivedAt  *time.Time
	OwnerId     *string
	Owner       *User
	PoolId      *uuid.UUID
	Pool        *TaskPool
}
