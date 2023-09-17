package model

import (
	"time"
)

type Post struct {
	Title     string
	Body      string
	OwnerID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
