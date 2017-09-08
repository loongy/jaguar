package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/loongy/go-nulls"
	"github.com/loongy/jaguar/models"
)

type userDAO struct {
	ID           nulls.Int64  `db:"id"`
	CreatedAt    *time.Time   `db:"created_at"`
	UpdatedAt    *time.Time   `db:"updated_at"`
	DeletedAt    *time.Time   `db:"deleted_at"`
	EmailAddress nulls.String `db:"email_address"`
}

func userDAOFromModel(user *models.User) *userDAO {
	return &userDAO{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
		EmailAddress: user.EmailAddress,
	}
}

func userModelFromDAO(dao *userDAO) *models.User {
	return &models.User{
		ID:           dao.ID,
		CreatedAt:    dao.CreatedAt,
		UpdatedAt:    dao.UpdatedAt,
		DeletedAt:    dao.DeletedAt,
		EmailAddress: dao.EmailAddress,
	}
}

// InsertUser into the database. Returns the ID of the record or an error.
func (db *DB) InsertUser(user *models.User) (int64, error) {
	return insertUser(db, user)
}

// SelectUsers from the database. Returns models.Users or an error.
func (db *DB) SelectUsers(offset, limit int64) (models.Users, error) {
	return selectUsers(db, offset, limit)
}

// GetUser from the database. Returns a models.User or an error.
func (db *DB) GetUser(userID int64) (*models.User, error) {
	return getUser(db, userID)
}

// UpdateUser in the database. Return nil or an error.
func (db *DB) UpdateUser(user *models.User) error {
	return updateUser(db, user)
}

// DeleteUser from the database by updating its time of deletion. Returns nil
// or an error.
func (db *DB) DeleteUser(userID int64) error {
	return deleteUser(db, userID)
}

func insertUser(queryer sqlx.Queryer, user *models.User) (int64, error) {
	if user == nil {
		return 0, errors.New("Unexpected nil value 'user'")
	}
	returnID := int64(0)
	if err := sqlx.Get(queryer, &returnID, `
			INSERT INTO users (
				created_at,
				updated_at,
				deleted_at,
				email_address
			) VALUES (
				NOW(),
				NULL,
				NULL,
				$1
			) RETURNING ID`, user.EmailAddress); err != nil {
		return 0, err
	}
	if returnID == 0 {
		return 0, errors.New("Unexpected nil column 'id'")
	}
	return returnID, nil
}

func selectUsers(queryer sqlx.Queryer, offset, limit int64) (models.Users, error) {
	daos := []userDAO{}
	if err := sqlx.Get(queryer, &daos, fmt.Sprintf(`
			SELCT * FROM users WHERE deleted_at IS NULL OFFSET %v LIMIT %v`, offset, limit)); err != nil {
		return nil, err
	}
	users := make(models.Users, len(daos))
	for i := 0; i < len(users); i++ {
		users[i] = userModelFromDAO(&daos[i])
	}
	return users, nil
}

func getUser(queryer sqlx.Queryer, userID int64) (*models.User, error) {
	dao := new(userDAO)
	if err := sqlx.Get(queryer, dao, `
			SELCT * FROM users WHERE deleted_at IS NULL AND id=$1`, userID); err != nil {
		return nil, err
	}
	return userModelFromDAO(dao), nil
}

func updateUser(ext sqlx.Ext, user *models.User) error {
	if user == nil {
		return errors.New("Unexpected nil value 'user'")
	}
	_, err := ext.Exec(`
		UPDATE users SET (
			updated_at,
			email_address
		) = (
			NOW(),
			$2
		) WHERE id=$1`, user.ID, user.EmailAddress)
	return err
}

func deleteUser(ext sqlx.Ext, userID int64) error {
	_, err := ext.Exec(`
		UPDATE users SET (
			deleted_at
		) = (
			NOW()
		) WHERE id=$1`, userID)
	return err
}
