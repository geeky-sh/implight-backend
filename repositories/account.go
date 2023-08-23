package repositories

import (
	"context"
	"errors"
	"implight-backend/domain"
	"implight-backend/utils"

	"github.com/jackc/pgx/v5"
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
	res := domain.User{}
	var id int

	sql := `
	INSERT INTO users (name, email, picture)
	VALUES ($1, $2, $3) RETURNING id`

	if err := r.db.QueryRow(ctx, sql, req.Name, req.Email, req.Picture).Scan(&id); err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	res = domain.User(req)
	res.ID = id

	return res, nil
}
func (r *accountRepository) GetByEmail(ctx context.Context, email string) (domain.User, utils.AppErr) {
	res := domain.User{}

	sql := `
	SELECT id, name, email, picture, created_at, updated_at
	from users where email=$1`
	if err := r.db.QueryRow(ctx, sql, email).Scan(&res.ID, &res.Name, &res.Email, &res.Picture, &res.CreatedAt, &res.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}

		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}
	return res, nil
}

type tokenRepository struct {
	table string
	db    *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool) domain.TokenRepository {
	return &tokenRepository{db: db, table: "tokens"}
}

func (r *tokenRepository) Create(ctx context.Context, req domain.AccessToken) (domain.AccessToken, utils.AppErr) {
	res := domain.AccessToken{}
	var idToken string

	sql := `
	INSERT INTO tokens (id_token, issued_at, expires_at, user_id)
	VALUES (gen_random_uuid(), $1, $2, $3) RETURNING id_token`
	if err := r.db.QueryRow(ctx, sql, req.IssuedAt, req.ExpiresAt, req.UserID).Scan(&idToken); err != nil {
		if err != nil {
			return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
		}
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	res = domain.AccessToken(req)
	res.IDToken = idToken

	return res, nil
}

func (r *tokenRepository) Get(ctx context.Context, tk string) (domain.AccessToken, utils.AppErr) {
	res := domain.AccessToken{}

	sql := `
	SELECT id_token, issued_at, expires_at, user_id
	FROM tokens WHERE id_token=$1`
	if err := r.db.QueryRow(ctx, sql, tk).Scan(&res.IDToken, &res.IssuedAt, &res.ExpiresAt, &res.UserID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return res, utils.NewAppErr(err.Error(), utils.ERR_OBJ_NOT_FOUND)
		}

		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	return res, nil
}
