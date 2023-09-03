package repositories

import (
	"context"
	"implight-backend/domain"
	"implight-backend/utils"

	"github.com/uptrace/bun"
)

type accountRepository struct {
	db *bun.DB
}

func NewAccountRepository(db *bun.DB) domain.AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Create(ctx context.Context, req domain.User) (domain.User, utils.AppErr) {
	res := domain.User{}
	var id int

	_, err := r.db.NewInsert().Model(&req).Returning("id").Exec(ctx, &id)
	if err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	res = domain.User(req)
	res.ID = id
	return res, nil
}

func (r *accountRepository) GetByEmail(ctx context.Context, email string) (domain.User, utils.AppErr) {
	res := domain.User{}

	if err := r.db.NewSelect().Model(&res).Where("email = ?", email).Limit(1).Scan(ctx); err != nil {
		if err.Error() == utils.ERR_MSG_SQL_NOT_FOUND {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	return res, nil
}

type tokenRepository struct {
	db *bun.DB
}

func NewTokenRepository(db *bun.DB) domain.TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(ctx context.Context, req domain.AccessToken) (domain.AccessToken, utils.AppErr) {
	res := domain.AccessToken{}

	_, err := r.db.NewInsert().Model(&req).Exec(ctx)
	if err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	res = domain.AccessToken(req)
	return res, nil
}

func (r *tokenRepository) Get(ctx context.Context, tk string) (domain.AccessToken, utils.AppErr) {
	res := domain.AccessToken{}

	if err := r.db.NewSelect().Model(&res).Where("uuid = ?", tk).Scan(ctx); err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	return res, nil
}
