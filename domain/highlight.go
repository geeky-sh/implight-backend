package domain

import (
	"context"
	"implight-backend/utils"
	"time"
)

// ---- schema ----

type Highlight struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	Text      string
	URL       string
}

type ListHighlight struct {
	Page   int
	Limit  int
	URL    string
	UserID int
}

// ---- end ----

// ---- request ----

type CreateHighlightReq struct {
	Text string `json:"text" validate:"required"`
	URL  string `json:"url" validate:"required,url"`
}

type ListHighlightsReq struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	URL   string `json:"url"`
}

// ---- end ----

// ---- schema ----

type GetHighlightRes struct {
	ID        int       `json:"id"`
	UserID    int       `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	URL       string    `json:"url"`
}

type ListHighlightRes struct {
	Count   int `json:"total_count"`
	Results []GetHighlightRes
}

// ---- end ----

// ---- interfaces ----

type HighlightUsecase interface {
	Create(ctx context.Context, by int, req CreateHighlightReq) (GetHighlightRes, utils.AppErr)
	List(ctx context.Context, by int, req ListHighlightsReq) (ListHighlightRes, utils.AppErr)
}

type HighlightRepository interface {
	Create(ctx context.Context, req Highlight) (Highlight, utils.AppErr)
	List(ctx context.Context, req ListHighlight) (int, []Highlight, utils.AppErr)
}

// ---- end ----

// ---- converters ----

func (r CreateHighlightReq) ToDB(userID int) Highlight {
	return Highlight{UserID: userID, CreatedAt: time.Now(), Text: r.Text, URL: r.URL}
}

func (r Highlight) ToRes() GetHighlightRes {
	return GetHighlightRes(r)
}

func (r ListHighlightsReq) ToDB(userID int) ListHighlight {
	return ListHighlight{Page: r.Page, Limit: r.Limit, URL: r.URL, UserID: userID}
}
