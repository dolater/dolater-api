package model

import "time"

type Notification struct {
	Title     string
	Body      string
	URL       string
	CreatedAt time.Time
}
