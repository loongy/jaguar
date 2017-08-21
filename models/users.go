package models

import (
	"time"

	"github.com/loongy/jaguar-template/nulls"
)

type User struct {
	CreatedAt    *time.Time   `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at"`
	ID           nulls.Int64  `json:"id"`
	EmailAddress nulls.String `json:"email_address"`
}

type Users []*User
