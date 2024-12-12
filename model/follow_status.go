package model

import "time"

type FollowStatus struct {
	FromId      string `gorm:"primaryKey"`
	From        User
	ToId        string `gorm:"primaryKey"`
	To          User
	RequestedAt time.Time
	ApprovedAt  *time.Time
}
