package actions

import "github.com/loongy/jaguar-template/models"

func CreateToken(ctx Context, token *models.Token) (*models.Token, error) {
	tokenID, err := ctx.Store.InsertToken(token)
	if err != nil {
		return nil, err
	}
	return GetToken(ctx, tokenID)
}

func GetTokens(ctx Context, offset, limit int64) (models.Tokens, error) {
	return ctx.Store.SelectTokens(offset, limit)
}

func GetToken(ctx Context, tokenID int64) (*models.Token, error) {
	return ctx.Store.GetToken(tokenID)
}

func DeleteToken(ctx Context, tokenID int64) error {
	return ctx.Store.DeleteToken(tokenID)
}
