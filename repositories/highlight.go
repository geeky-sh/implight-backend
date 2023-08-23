package repositories

import (
	"context"
	"implight-backend/domain"
	"implight-backend/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type highlightRepository struct {
	db    *pgxpool.Pool
	table string
}

func NewHighlightRepository(db *pgxpool.Pool) domain.HighlightRepository {
	return &highlightRepository{db: db, table: "highlights"}
}

func (r *highlightRepository) Create(ctx context.Context, req domain.Highlight) (domain.Highlight, utils.AppErr) {
	res := domain.Highlight{}
	var id int

	sql := `
	INSERT INTO highlights (user_id, created_at, text, url)
	VALUES ($1, $1, $3, $4) RETURNING id`

	if err := r.db.QueryRow(ctx, sql, req.UserID, req.CreatedAt, req.Text, req.URL).Scan(&id); err != nil {
		return res, utils.NewAppErr(err.Error(), utils.ERR_UNKNOWN)
	}

	res = domain.Highlight(req)
	res.ID = id

	return res, nil
}

func (r *highlightRepository) List(ctx context.Context, req domain.ListHighlight) (int, []domain.Highlight, utils.AppErr) {
	return 0, []domain.Highlight{}, nil
}
