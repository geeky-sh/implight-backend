package domain

import (
	"context"
	"implight-backend/utils"
	"time"

	"github.com/uptrace/bun"
	"google.golang.org/api/idtoken"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int       `bun:"id,pk,autoincrement"`
	Email         string    `bun:"email,notnull"`
	Name          string    `bun:"name"`
	Picture       string    `bun:"picture"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}

type AccessToken struct {
	bun.BaseModel `bun:"table:tokens"`
	UUID          string    `json:"uuid" bun:"uuid,pk"`
	IDToken       string    `json:"id_token" bun:"id_token,notnull"`
	IssuedAt      time.Time `json:"issued_at" bun:"issued_at,notnull"`
	ExpiresAt     time.Time `json:"expires_at" bun:"expires_at,notnull"`
	UserID        int       `json:"-" bun:"user_id,notnull"`
	User          User      `json:"-" bun:"rel:belongs-to,join:user_id=id"`
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
