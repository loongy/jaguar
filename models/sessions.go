package models

import (
	"time"

	"github.com/loongy/jaguar/nulls"
)

type Session struct {
	CreatedAt *time.Time  `json:"created_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
	ID        nulls.Int64 `json:"id"`
	UserID    nulls.Int64 `json:"user_id"`
}

type Sessions []*Session
