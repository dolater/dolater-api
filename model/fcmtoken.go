package model

import (
	"time"
)

type FCMToken struct {
	RegistrationToken string `gorm:"primaryKey"`
	UserId            string
	Timestamp         time.Time
}
