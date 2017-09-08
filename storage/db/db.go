package db

import (
	"github.com/jmoiron/sqlx"
)

// DB is a connection to an SQL database. It implements the DataStore
// interface.
type DB struct {
	*sqlx.DB
}

// NewDB returns a new DB, or an error. It will attempt a ping to guarantee
// that a connection has been established.
func NewDB(driver, url string) (*DB, error) {
	db, err := sqlx.Open(driver, url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
