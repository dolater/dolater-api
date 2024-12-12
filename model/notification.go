package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserId    string
	Title     string
	Body      string
	URL       string
	CreatedAt time.Time
}
