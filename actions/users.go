package actions

import "github.com/loongy/jaguar-template/models"

func CreateUser(ctx Context, user *models.User) (*models.User, error) {
	return ctx.Store.InsertUser(user)
}

func GetUsers(ctx Context, offset, limit int64) (models.Users, error) {
	return ctx.Store.SelectUsers(offset, limit)
}

func GetUser(ctx Context, userID int64) (*models.User, error) {
	return ctx.Store.GetUser(userID)
}
