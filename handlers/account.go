package handlers

import (
	"implight-backend/domain"
	"implight-backend/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/api/idtoken"
)

type accountHandler struct {
	db *pgxpool.Pool
	uc domain.AccountUsecase
}

func NewAccountHandler(db *pgxpool.Pool, uc domain.AccountUsecase) accountHandler {
	return accountHandler{db: db, uc: uc}
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
	token := r.FormValue("credential")

	payload, err := idtoken.Validate(ctx, token, "79775009631-ba2k1s28mru34jo8ephd50h15ouink6o.apps.googleusercontent.com")
	if err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, "Payload Validation Failed")
		return
	}

	res, aerr := h.uc.LogIn(ctx, token, payload)
	if aerr != nil {
		utils.WriteAppErrRes(w, aerr)
		return
	}

	utils.WriteRes(w, http.StatusOK, res)
}
