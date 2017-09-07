package actions

import (
	"github.com/loongy/jaguar/storage"
	"github.com/loongy/jaguar/storage/db"
)

//
type Context struct {
	DataStore storage.DataStore
	FileStore storage.FileStore
}

//
func NewContext(dbDriver, dbURL string) (*Context, error) {
	db, err := db.NewDB(dbDriver, dbURL)
	if err != nil {
		return nil, err
	}
	return &Context{DataStore: db}, nil
}
