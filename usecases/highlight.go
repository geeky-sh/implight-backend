package usecases

import (
	"context"
	"implight-backend/domain"
	"implight-backend/utils"
)

type highlightUsecase struct {
	repo domain.HighlightRepository
}

func NewHighlightUsecase(r domain.HighlightRepository) domain.HighlightUsecase {
	return &highlightUsecase{repo: r}
}

func (u *highlightUsecase) Create(ctx context.Context, by int, req domain.CreateHighlightReq) (domain.GetHighlightRes, utils.AppErr) {
	res := domain.GetHighlightRes{}
	dbReq := req.ToDB(by)

	dbRes, err := u.repo.Create(ctx, dbReq)
	if err != nil {
		return res, err
	}
	return dbRes.ToRes(), nil
}

func (u *highlightUsecase) List(ctx context.Context, by int, req domain.ListHighlightsReq) (domain.ListHighlightRes, utils.AppErr) {
	res := domain.ListHighlightRes{}
	dbReq := req.ToDB(by)

	cnt, dbRes, err := u.repo.List(ctx, dbReq)
	if err != nil {
		return res, err
	}
	res.Count = cnt
	for _, r := range dbRes {
		res.Results = append(res.Results, r.ToRes())
	}

	return res, nil
}
