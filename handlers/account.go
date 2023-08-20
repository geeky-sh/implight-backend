package handlers

import (
	"fmt"
	"implight-backend/utils"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/api/idtoken"
)

type accountHandler struct {
	db *pgxpool.Pool
}

func NewAccountHandler(db *pgxpool.Pool) accountHandler {
	return accountHandler{db: db}
}

func (h *accountHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/login", h.LoginViaGooglePage)
	r.Post("/callback", h.Callback)
	return r
}

func (h *accountHandler) LoginViaGooglePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/Users/aash/projects/personal/implight-backend/static/login-google.html")
}

func (h *accountHandler) Callback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := r.ParseForm(); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	payload, err := idtoken.Validate(ctx, r.FormValue("credential"), "79775009631-ba2k1s28mru34jo8ephd50h15ouink6o.apps.googleusercontent.com")
	if err != nil {
		panic(err)
	}

	ea := time.Unix(payload.Expires, 0)
	fmt.Println(ea)

	ia := time.Unix(payload.IssuedAt, 0)
	fmt.Println(ia)

	utils.WriteRes(w, http.StatusOK, payload)
}
