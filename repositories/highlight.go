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
	return domain.Highlight{}, nil
}

func (r *highlightRepository) List(ctx context.Context, req domain.ListHighlight) (int, []domain.Highlight, utils.AppErr) {
	return 0, []domain.Highlight{}, nil
}
