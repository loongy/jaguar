package models

import (
	"time"

	"github.com/loongy/jaguar-template/nulls"
)

type Token struct {
	CreatedAt *time.Time  `json:"created_at"`
	ID        nulls.Int64 `json:"id"`
	UserID    nulls.Int64 `json:"user_id"`
}

type Tokens []*Token
