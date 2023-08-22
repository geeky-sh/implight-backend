package handlers

import (
	"implight-backend/domain"
	"implight-backend/middlewares"
	"implight-backend/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type highlightHandler struct {
	uc domain.HighlightUsecase
	m  *middlewares.Middleware
	v  *validator.Validate
}

func NewHighlightHandler(uc domain.HighlightUsecase, m *middlewares.Middleware, v *validator.Validate) highlightHandler {
	return highlightHandler{uc: uc, m: m, v: v}
}

func (h *highlightHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(h.m.Authenticate)
	r.Get("/", h.List)
	r.Post("/", h.Create)
	return r
}

func (h *highlightHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := domain.ListHighlightsReq{}

	userID, aerr := middlewares.UserID(ctx)
	if aerr != nil {
		utils.WriteAppErrRes(w, aerr)
		return
	}

	if err := utils.JSNDecode(r, &req); err != nil {
		utils.WriteAppErrRes(w, err)
		return
	}

	res, aerr := h.uc.List(ctx, userID, req)
	if aerr != nil {
		utils.WriteAppErrRes(w, aerr)
		return
	}

	utils.WriteRes(w, http.StatusOK, res)
}

func (h *highlightHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := domain.CreateHighlightReq{}

	userID, aerr := middlewares.UserID(ctx)
	if aerr != nil {
		utils.WriteAppErrRes(w, aerr)
		return
	}

	if err := utils.JSNDecode(r, &req); err != nil {
		utils.WriteAppErrRes(w, err)
		return
	}

	if err := h.v.Struct(&req); err != nil {
		utils.WriteMsgRes(w, http.StatusBadRequest, err.Error())
		return
	}

	res, aerr := h.uc.Create(ctx, userID, req)
	if aerr != nil {
		utils.WriteAppErrRes(w, aerr)
		return
	}

	utils.WriteRes(w, http.StatusOK, res)
}
