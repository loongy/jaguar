package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/loongy/go-nulls"
	"github.com/loongy/jaguar/models"
)

type sessionDao struct {
	ID        nulls.Int64 `db:"id"`
	CreatedAt *time.Time  `db:"created_at"`
	DeletedAt *time.Time  `db:"deleted_at"`
	UserID    nulls.Int64 `db:"user_id"`
}

func sessionDAOFromModel(session *models.Session) *sessionDao {
	return &sessionDao{
		ID:        session.ID,
		CreatedAt: session.CreatedAt,
		DeletedAt: session.DeletedAt,
		UserID:    session.UserID,
	}
}

func sessionModelFromDAO(dao *sessionDao) *models.Session {
	return &models.Session{
		ID:        dao.ID,
		CreatedAt: dao.CreatedAt,
		DeletedAt: dao.DeletedAt,
		UserID:    dao.UserID,
	}
}

// InsertSession into the database. Returns the ID of the record or an error.
func (db *DB) InsertSession(session *models.Session) (int64, error) {
	return insertSession(db, session)
}

// SelectSessions from the database. Returns models.Sessions or an error.
func (db *DB) SelectSessions(offset, limit int64) (models.Sessions, error) {
	return selectSessions(db, offset, limit)
}

// GetSession from the database. Returns a models.Session or an error.
func (db *DB) GetSession(sessionID int64) (*models.Session, error) {
	return getSession(db, sessionID)
}

// DeleteSession from the database by updating its time of deletion. Returns
// nil or an error.
func (db *DB) DeleteSession(sessionID int64) error {
	return deleteSession(db, sessionID)
}

func insertSession(queryer sqlx.Queryer, session *models.Session) (int64, error) {
	if session == nil {
		return 0, errors.New("Unexpected nil value 'session'")
	}
	returnID := int64(0)
	if err := sqlx.Get(queryer, &returnID, `
			INSERT INTO sessions (
				created_at,
				deleted_at,
				user_id
			) VALUES (
				NOW(),
				NULL,
				$1
			) RETURNING ID`, session.UserID); err != nil {
		return 0, err
	}
	if returnID == 0 {
		return 0, errors.New("Unexpected nil column 'id'")
	}
	return returnID, nil
}

func selectSessions(queryer sqlx.Queryer, offset, limit int64) (models.Sessions, error) {
	daos := []sessionDao{}
	if err := sqlx.Get(queryer, &daos, fmt.Sprintf(`
			SELECT * FROM sessions WHERE deleted_at IS NULL OFFSET %v LIMIT %v`, offset, limit)); err != nil {
		return nil, err
	}
	sessions := make(models.Sessions, len(daos))
	for i := 0; i < len(sessions); i++ {
		sessions[i] = sessionModelFromDAO(&daos[i])
	}
	return sessions, nil
}

func getSession(queryer sqlx.Queryer, sessionID int64) (*models.Session, error) {
	dao := new(sessionDao)
	if err := sqlx.Get(queryer, dao, `
			SELECT * FROM sessions WHERE deleted_at IS NULL AND id=$1`, sessionID); err != nil {
		return nil, err
	}
	return sessionModelFromDAO(dao), nil
}

func deleteSession(ext sqlx.Ext, sessionID int64) error {
	_, err := ext.Exec(`
		UPDATE sessions SET (
			deleted_at
		) = (
			NOW()
		) WHERE id=$1`, sessionID)
	return err
}
