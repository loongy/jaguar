package storage

import "github.com/loongy/jaguar-template/models"

type Store interface {
	InsertUser(user *models.User) (*models.User, error)
	SelectUsers(offset, limit int64) (models.Users, error)
	GetUser(userID int64) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}
