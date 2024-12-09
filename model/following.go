package model

import (
	"time"
)

type Following struct {
	UserId       string `gorm:"primaryKey"`
	User         User
	TargetUserId string `gorm:"primaryKey"`
	TargetUser   User
	RequestedAt  time.Time
	ApprovedAt   *time.Time
	IsFollowed   bool
}
