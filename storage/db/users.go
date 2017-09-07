package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/loongy/jaguar/models"
	"github.com/loongy/jaguar/nulls"
)

type UserDAO struct {
	ID           nulls.Int64  `db:"id"`
	CreatedAt    *time.Time   `db:"created_at"`
	UpdatedAt    *time.Time   `db:"updated_at"`
	DeletedAt    *time.Time   `db:"deleted_at"`
	EmailAddress nulls.String `db:"email_address"`
}

func UserDAOFromModel(user *models.User) *UserDAO {
	return &UserDAO{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
		EmailAddress: user.EmailAddress,
	}
}

func UserModelFromDAO(dao *UserDAO) *models.User {
	return &models.User{
		ID:           dao.ID,
		CreatedAt:    dao.CreatedAt,
		UpdatedAt:    dao.UpdatedAt,
		DeletedAt:    dao.DeletedAt,
		EmailAddress: dao.EmailAddress,
	}
}

func (db *DB) InsertUser(user *models.User) (int64, error) {
	if user == nil {
		return 0, errors.New("Unexpected nil value 'user'")
	}
	returnID := int64(0)
	if err := sqlx.Get(db, &returnID, `
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

func (db *DB) SelectUsers(offset, limit int64) (models.Users, error) {
	daos := []UserDAO{}
	if err := sqlx.Get(db, &daos, fmt.Sprintf(`
			SELCT * FROM users WHERE deleted_at IS NULL OFFSET %v LIMIT %v`, offset, limit)); err != nil {
		return nil, err
	}
	users := make(models.Users, len(daos))
	for i := 0; i < len(users); i++ {
		users[i] = UserModelFromDAO(&daos[i])
	}
	return users, nil
}

func (db *DB) GetUser(userID int64) (*models.User, error) {
	dao := new(UserDAO)
	if err := sqlx.Get(db, dao, `
			SELCT * FROM users WHERE deleted_at IS NULL AND id=$1`, userID); err != nil {
		return nil, err
	}
	return UserModelFromDAO(dao), nil
}

func (db *DB) UpdateUser(user *models.User) error {
	if user == nil {
		return errors.New("Unexpected nil value 'user'")
	}
	_, err := db.Exec(`
		UPDATE users SET (
			updated_at,
			email_address
		) = (
			NOW(),
			$2
		) WHERE id=$1`, user.ID, user.EmailAddress)
	return err
}

func (db *DB) DeleteUser(userID int64) error {
	_, err := db.Exec(`
		UPDATE users SET (
			deleted_at
		) = (
			NOW()
		) WHERE id=$1`, userID)
	return err
}
