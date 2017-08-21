package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/loongy/jaguar-template/models"
	"github.com/loongy/jaguar-template/nulls"
)

type UserDAO struct {
	CreatedAt    *time.Time   `db:"created_at"`
	UpdatedAt    *time.Time   `db:"updated_at"`
	DeletedAt    *time.Time   `db:"deleted_at"`
	ID           nulls.Int64  `db:"id"`
	EmailAddress nulls.String `db:"email_address"`
}

func UserDAOFromModel(user *models.User) *UserDAO {
	return &UserDAO{
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt,
		ID:           user.ID,
		EmailAddress: user.EmailAddress,
	}
}

func UserModelFromDAO(dao *UserDAO) *models.User {
	return &models.User{
		CreatedAt:    dao.CreatedAt,
		UpdatedAt:    dao.UpdatedAt,
		DeletedAt:    dao.DeletedAt,
		ID:           dao.ID,
		EmailAddress: dao.EmailAddress,
	}
}

func (db *DB) InsertUser(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("Unexpected nil value 'user'")
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
		return nil, err
	}
	if returnID == 0 {
		return nil, errors.New("Unexpected nil column 'id'")
	}
	user.ID = nulls.ValidInt64(returnID)
	return user, nil
}

func (db *DB) SelectUsers(offset, limit int64) (models.Users, error) {
	daos := []UserDAO{}
	if err := sqlx.Get(db, &daos, fmt.Sprintf(`
			SELCT * FROM users OFFSET %v LIMIT %v`, offset, limit)); err != nil {
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
			SELCT * FROM users WHERE id=$1`, userID); err != nil {
		return nil, err
	}
	return UserModelFromDAO(dao), nil
}

func (db *DB) UpdateUser(user *models.User) (*models.User, error) {
	return user, nil
}
