package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/loongy/jaguar/models"
	"github.com/loongy/go-nulls"
)

type SessionDAO struct {
	ID        nulls.Int64 `db:"id"`
	CreatedAt *time.Time  `db:"created_at"`
	DeletedAt *time.Time  `db:"deleted_at"`
	UserID    nulls.Int64 `db:"user_id"`
}

func SessionDAOFromModel(session *models.Session) *SessionDAO {
	return &SessionDAO{
		ID:        session.ID,
		CreatedAt: session.CreatedAt,
		DeletedAt: session.DeletedAt,
		UserID:    session.UserID,
	}
}

func SessionModelFromDAO(dao *SessionDAO) *models.Session {
	return &models.Session{
		ID:        dao.ID,
		CreatedAt: dao.CreatedAt,
		DeletedAt: dao.DeletedAt,
		UserID:    dao.UserID,
	}
}

func (db *DB) InsertSession(session *models.Session) (int64, error) {
	if session == nil {
		return 0, errors.New("Unexpected nil value 'session'")
	}
	returnID := int64(0)
	if err := sqlx.Get(db, &returnID, `
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

func (db *DB) SelectSessions(offset, limit int64) (models.Sessions, error) {
	daos := []SessionDAO{}
	if err := sqlx.Get(db, &daos, fmt.Sprintf(`
			SELECT * FROM sessions WHERE deleted_at IS NULL OFFSET %v LIMIT %v`, offset, limit)); err != nil {
		return nil, err
	}
	sessions := make(models.Sessions, len(daos))
	for i := 0; i < len(sessions); i++ {
		sessions[i] = SessionModelFromDAO(&daos[i])
	}
	return sessions, nil
}

func (db *DB) GetSession(sessionID int64) (*models.Session, error) {
	dao := new(SessionDAO)
	if err := sqlx.Get(db, dao, `
			SELECT * FROM sessions WHERE deleted_at IS NULL AND id=$1`, sessionID); err != nil {
		return nil, err
	}
	return SessionModelFromDAO(dao), nil
}

func (db *DB) DeleteSession(sessionID int64) error {
	_, err := db.Exec(`
		UPDATE sessions SET (
			deleted_at
		) = (
			NOW()
		) WHERE id=$1`, sessionID)
	return err
}
