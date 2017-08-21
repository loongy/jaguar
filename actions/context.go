package actions

import (
	"os"

	"github.com/loongy/jaguar/storage"
	"github.com/loongy/jaguar/storage/db"
)

type Context struct {
	Store storage.Store
}

func NewContext() (*Context, error) {
	db, err := db.NewPostgresDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &Context{Store: db}, nil
}
