package usecases

import (
	"context"
	"implight-backend/domain"
	"implight-backend/utils"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
)

type accountUsecase struct {
	arepo domain.AccountRepository
	trepo domain.TokenRepository
}

func NewAccountUsecase(ar domain.AccountRepository, tr domain.TokenRepository) domain.AccountUsecase {
	return &accountUsecase{arepo: ar, trepo: tr}
}

func (u *accountUsecase) LogIn(ctx context.Context, token string, pl *idtoken.Payload) (domain.AccessToken, utils.AppErr) {
	/*
		Steps:
		1. Check if the user is already present
		2. If yes, store the token received and return
		3. If no, create the user, store the token received and return
	*/
	res := domain.AccessToken{}

	user, err := u.arepo.GetByEmail(ctx, pl.Claims["email"].(string))
	if err != nil && err.ErrCode() != utils.ERR_OBJ_NOT_FOUND {
		return res, err
	} else if err != nil {
		// create user
		req := domain.User{
			Email: pl.Claims["email"].(string), Name: pl.Claims["name"].(string),
			CreatedAt: time.Now(), UpdatedAt: time.Now(), Picture: pl.Claims["picture"].(string)}
		user, err = u.arepo.Create(ctx, req)
		if err != nil {
			return res, err
		}
	}

	tkreq := domain.AccessToken{UserID: user.ID, IssuedAt: time.Unix(pl.IssuedAt, 0),
		ExpiresAt: time.Unix(pl.Expires, 0), IDToken: token, UUID: uuid.New().String()}
	res, err = u.trepo.Create(ctx, tkreq)
	if err != nil {
		return res, err
	}
	return res, err
}
