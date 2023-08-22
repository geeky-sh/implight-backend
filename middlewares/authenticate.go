package middlewares

import (
	"context"
	"implight-backend/domain"
	"implight-backend/utils"
	"net/http"
	"strings"
	"time"
)

type Middleware struct {
	tk domain.TokenRepository
}

func New(tk domain.TokenRepository) *Middleware {
	return &Middleware{tk: tk}
}

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		h := r.Header.Get("Authorization")

		lst := strings.Split(h, " ")
		if len(lst) != 2 {
			utils.WriteMsgRes(w, http.StatusBadRequest, "Incorrect Headers")
			return
		}

		token := lst[1]
		res, err := m.tk.Get(ctx, token)
		if err != nil {
			if err.ErrCode() == utils.ERR_OBJ_NOT_FOUND {
				utils.WriteMsgRes(w, http.StatusBadRequest, "token not found")
				return
			}
			utils.WriteAppErrRes(w, err)
			return
		}
		if res.ExpiresAt.Before(time.Now()) {
			utils.WriteMsgRes(w, http.StatusBadRequest, "token has expired")
			return
		}
		ctx = context.WithValue(ctx, "user_id", res.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserID(ctx context.Context) (int, utils.AppErr) {
	userID := ctx.Value("user_id")
	if userID == "" {
		return 0, utils.NewAppErr("Can't fetch user_id", utils.ERR_UNKNOWN)
	}
	return userID.(int), nil
}
