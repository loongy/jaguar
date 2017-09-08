package models

import (
	"time"

	"github.com/loongy/go-nulls"
)

type User struct {
	ID           nulls.Int64  `json:"id"`
	CreatedAt    *time.Time   `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at"`
	EmailAddress nulls.String `json:"email_address"`
}

type Users []*User
