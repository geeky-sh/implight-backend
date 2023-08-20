package domain

import (
	"context"
	"implight-backend/utils"
	"time"

	"google.golang.org/api/idtoken"
)

type User struct {
	ID        string
	Email     string
	Name      string
	Picture   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccessToken struct {
	IDToken   string    `json:"id_token"`
	IssuedAt  time.Time `json:"expires_at"`
	ExpiresAt time.Time `json:"issued_at"`
	UserID    string    `json:"-"`
}

type AccountUsecase interface {
	LogIn(ctx context.Context, token string, pl *idtoken.Payload) (AccessToken, utils.AppErr)
}

type AccountRepository interface {
	Create(ctx context.Context, req User) (User, utils.AppErr)
	GetByEmail(ctx context.Context, email string) (User, utils.AppErr)
}

type TokenRepository interface {
	Create(ctx context.Context, req AccessToken) (AccessToken, utils.AppErr)
	Get(ctx context.Context, tk string) (AccessToken, utils.AppErr)
}
