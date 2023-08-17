package handlers

import (
	"fmt"
	"implight-backend/domain"
	"implight-backend/utils"
	"net/http"

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
	req := domain.GoogleCallback{}

	if err := r.ParseForm(); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	req.Credential = r.FormValue("credential")
	req.GCSRFToken = r.FormValue("g_csrf_token")

	payload, err := idtoken.Validate(ctx, req.Credential, "79775009631-ba2k1s28mru34jo8ephd50h15ouink6o.apps.googleusercontent.com")
	if err != nil {
		panic(err)
	}
	fmt.Print(payload.Claims)

	utils.WriteMsgRes(w, http.StatusOK, "Success")
}
