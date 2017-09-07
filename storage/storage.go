package storage

import (
	"io"

	"github.com/loongy/jaguar/models"
)

type DataStore interface {
	SessionStore
	UserStore
}

type FileStore interface {
	SaveImage(filename string, reader io.Reader) (*models.Image, error)
}

type SessionStore interface {
	InsertSession(session *models.Session) (int64, error)
	SelectSessions(offset, limit int64) (models.Sessions, error)
	GetSession(sessionID int64) (*models.Session, error)
	DeleteSession(sessionID int64) error
}

type UserStore interface {
	InsertUser(user *models.User) (int64, error)
	SelectUsers(offset, limit int64) (models.Users, error)
	GetUser(userID int64) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID int64) error
}
