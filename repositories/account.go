package repositories

import (
	"context"
	"implight-backend/domain"
	"implight-backend/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	table string
	db    *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) domain.AccountRepository {
	return &accountRepository{"users", db}
}

func (r *accountRepository) Create(ctx context.Context, req domain.User) (domain.User, utils.AppErr) {
	return domain.User{}, nil
}
func (r *accountRepository) GetByEmail(ctx context.Context, email string) (domain.User, utils.AppErr) {
	return domain.User{}, nil
}

type tokenRepository struct {
	table string
	db    *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool) domain.TokenRepository {
	return &tokenRepository{db: db, table: "tokens"}
}

func (r *tokenRepository) Create(ctx context.Context, req domain.AccessToken) (domain.AccessToken, utils.AppErr) {
	return domain.AccessToken{}, nil
}

func (r *tokenRepository) Get(ctx context.Context, tk string) (domain.AccessToken, utils.AppErr) {
	return domain.AccessToken{}, nil
}
