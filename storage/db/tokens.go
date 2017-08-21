package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/loongy/jaguar-template/models"
	"github.com/loongy/jaguar-template/nulls"
)

type TokenDAO struct {
	CreatedAt *time.Time  `db:"created_at"`
	ID        nulls.Int64 `db:"id"`
	UserID    nulls.Int64 `db:"user_id"`
}

func TokenDAOFromModel(token *models.Token) *TokenDAO {
	return &TokenDAO{
		CreatedAt: token.CreatedAt,
		ID:        token.ID,
		UserID:    token.UserID,
	}
}

func TokenModelFromDAO(dao *TokenDAO) *models.Token {
	return &models.Token{
		CreatedAt: dao.CreatedAt,
		ID:        dao.ID,
		UserID:    dao.UserID,
	}
}

func (db *DB) InsertToken(token *models.Token) (int64, error) {
	if token == nil {
		return 0, errors.New("Unexpected nil value 'token'")
	}
	returnID := int64(0)
	if err := sqlx.Get(db, &returnID, `
			INSERT INTO tokens (
				created_at,
				user_id
			) VALUES (
				NOW(),
				$1
			) RETURNING ID`, token.UserID); err != nil {
		return 0, err
	}
	if returnID == 0 {
		return 0, errors.New("Unexpected nil column 'id'")
	}
	return returnID, nil
}

func (db *DB) SelectTokens(offset, limit int64) (models.Tokens, error) {
	daos := []TokenDAO{}
	if err := sqlx.Get(db, &daos, fmt.Sprintf(`
			SELCT * FROM tokens OFFSET %v LIMIT %v`, offset, limit)); err != nil {
		return nil, err
	}
	tokens := make(models.Tokens, len(daos))
	for i := 0; i < len(tokens); i++ {
		tokens[i] = TokenModelFromDAO(&daos[i])
	}
	return tokens, nil
}

func (db *DB) GetToken(tokenID int64) (*models.Token, error) {
	dao := new(TokenDAO)
	if err := sqlx.Get(db, dao, `
			SELCT * FROM tokens WHERE id=$1`, tokenID); err != nil {
		return nil, err
	}
	return TokenModelFromDAO(dao), nil
}

func (db *DB) DeleteToken(tokenID int64) error {
	_, err := db.Exec(`
		DELETE FROM tokens WHERE id=$1`, tokenID)
	return err
}
