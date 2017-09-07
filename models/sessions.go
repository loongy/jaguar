package models

import (
	"time"

	"github.com/loongy/jaguar/nulls"
)

type Session struct {
	ID        nulls.Int64 `json:"id"`
	CreatedAt *time.Time  `json:"created_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
	UserID    nulls.Int64 `json:"user_id"`
}

type Sessions []*Session
