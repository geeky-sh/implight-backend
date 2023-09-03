package main

import (
	"database/sql"
	"implight-backend/handlers"
	"implight-backend/middlewares"
	"implight-backend/repositories"
	"implight-backend/usecases"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Unable to load .env %v\n", err)
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DB_URL"))))
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	v := validator.New()
	tr := repositories.NewTokenRepository(db)
	m := middlewares.New(tr)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(3 * time.Second))

	hh := handlers.NewHealthHandler(db)
	r.Mount("/metrics", hh.Routes())

	ar := repositories.NewAccountRepository(db)
	auc := usecases.NewAccountUsecase(ar, tr)
	ah := handlers.NewAccountHandler(db, auc)
	r.Mount("/account", ah.Routes())

	hr := repositories.NewHighlightRepository(db)
	huc := usecases.NewHighlightUsecase(hr)
	hhl := handlers.NewHighlightHandler(huc, m, v)
	r.Mount("/highlights", hhl.Routes())

	http.ListenAndServe(":3000", r)
}
