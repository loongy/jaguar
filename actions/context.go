package actions

import (
	"github.com/dappbase/api/models"
	"github.com/loongy/jaguar/storage"
)

// Context holds information that is necessary for actions to complete.
type Context struct {
	// DataStore is used for writing and reading to and from persistent storage.
	storage.DataStore
	// FileStore is used for writing and reading to and from a file system.
	storage.FileStore

	// Me is nil if there is no authenticated user
	Me *models.User

	// Values is a map for holding arbitrary key-value pairs.
	Values map[interface{}]interface{}
}

// NewContext returns a new context.
func NewContext(dataStore storage.DataStore, fileStore storage.FileStore) Context {
	return Context{
		DataStore: dataStore,
		FileStore: fileStore,
		Values:    make(map[interface{}]interface{}),
	}
}
