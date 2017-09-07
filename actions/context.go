package actions

import (
	"github.com/loongy/jaguar/storage"
	"github.com/loongy/jaguar/storage/db"
)

//
type Context struct {
	Store storage.Store
}

//
func NewContext(dbURL string) (*Context, error) {
	db, err := db.NewPostgresDB(dbURL)
	if err != nil {
		return nil, err
	}
	return &Context{Store: db}, nil
}
