package model

import "time"

type FollowStatus struct {
	FromId      string
	From        User
	ToId        string
	To          User
	RequestedAt time.Time
	ApprovedAt  *time.Time
}
