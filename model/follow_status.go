package model

import "time"

type FollowStatus struct {
	From        User
	To          User
	RequestedAt time.Time
	ApprovedAt  *time.Time
}
