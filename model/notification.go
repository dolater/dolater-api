package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	Id        uuid.UUID
	UserId    string
	Title     string
	Body      string
	URL       string
	CreatedAt time.Time
}
