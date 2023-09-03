package handlers

import (
	"implight-backend/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)

type healthHandler struct {
	db *bun.DB
}

func NewHealthHandler(db *bun.DB) healthHandler {
	return healthHandler{db}
}

func (h *healthHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/health", h.HealthCheck)
	return r
}

func (h *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := h.db.Ping(); err != nil {
		utils.WriteMsgRes(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteMsgRes(w, http.StatusOK, "OK")
}
