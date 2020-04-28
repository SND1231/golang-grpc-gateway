package model

import (
	"time"
)

type Post struct {
	ID        int32
	Tittle    string
	Content   string
	PhotoUrl  string
	UserId    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
