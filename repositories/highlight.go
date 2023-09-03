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
	res := []domain.Highlight{}
	offset := (req.Page - 1) * req.Limit

	q := r.db.NewSelect().Model(&res).Limit(req.Limit)
	if req.URL != "" {
		q = q.Where("url = ?", req.URL)
	}
	if req.UserID != 0 {
		q = q.Where("user_id = ?", req.UserID)
	}
	if offset > 0 {
		q = q.Offset(offset)
	}
	count, err := q.ScanAndCount(ctx)
	if err != nil {
		return 0, res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	return count, res, nil
}
