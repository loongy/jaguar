package storage

import (
	"io"

	"github.com/loongy/jaguar/models"
)

type Store interface {
	TokenStore
	UserStore
}

type FileStore interface {
	SaveImage(filename string, reader io.Reader) (*models.Image, error)
}

type TokenStore interface {
	InsertToken(token *models.Token) (int64, error)
	SelectTokens(offset, limit int64) (models.Tokens, error)
	GetToken(tokenID int64) (*models.Token, error)
	DeleteToken(tokenID int64) error
}

type UserStore interface {
	InsertUser(user *models.User) (int64, error)
	SelectUsers(offset, limit int64) (models.Users, error)
	GetUser(userID int64) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID int64) error
}
