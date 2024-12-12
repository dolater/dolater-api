package model

import "time"

type FollowStatus struct {
	FromUserId  string `gorm:"column:from_user_id"`
	From        User   `gorm:"foreignkey:FromUserId"`
	ToUserId    string `gorm:"column:to_user_id"`
	To          User   `gorm:"foreignkey:ToUserId"`
	RequestedAt time.Time
	ApprovedAt  *time.Time
}
