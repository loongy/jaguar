package storage

import (
	"io"

	"github.com/loongy/jaguar-template/models"
)

type Store interface {
	UserStore
}

type FileStore interface {
	SaveImage(filename string, reader io.Reader) (*models.Image, error)
}

type UserStore interface {
	InsertUser(user *models.User) (*models.User, error)
	SelectUsers(offset, limit int64) (models.Users, error)
	GetUser(userID int64) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}
