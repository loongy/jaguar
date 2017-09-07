package actions

import "github.com/loongy/jaguar/models"

func CreateSession(ctx Context, session *models.Session) (*models.Session, error) {
	sessionID, err := ctx.DataStore.InsertSession(session)
	if err != nil {
		return nil, err
	}
	return GetSession(ctx, sessionID)
}

func GetSessions(ctx Context, offset, limit int64) (models.Sessions, error) {
	return ctx.DataStore.SelectSessions(offset, limit)
}

func GetSession(ctx Context, sessionID int64) (*models.Session, error) {
	return ctx.DataStore.GetSession(sessionID)
}

func DeleteSession(ctx Context, sessionID int64) error {
	return ctx.DataStore.DeleteSession(sessionID)
}
