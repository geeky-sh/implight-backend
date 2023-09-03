package repositories

import (
	"context"
	"implight-backend/domain"
	"implight-backend/utils"

	"github.com/uptrace/bun"
)

type highlightRepository struct {
	db *bun.DB
}

func NewHighlightRepository(db *bun.DB) domain.HighlightRepository {
	return &highlightRepository{db: db}
}

func (r *highlightRepository) Create(ctx context.Context, req domain.Highlight) (domain.Highlight, utils.AppErr) {
	res := domain.Highlight{}
	var id int

	_, err := r.db.NewInsert().Model(&req).Returning("id").Exec(ctx, &id)
	if err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	res = domain.Highlight(req)
	res.ID = id

	return res, nil
}

func (r *highlightRepository) List(ctx context.Context, req domain.ListHighlight) (int, []domain.Highlight, utils.AppErr) {
	return 0, []domain.Highlight{}, nil
}
