package main

import (
	"context"
	"implight-backend/handlers"
	"implight-backend/repositories"
	"implight-backend/usecases"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Unable to load .env %v\n", err)
	}

	db, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// v := validator.New()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(3 * time.Second))

	hh := handlers.NewHealthHandler(db)
	r.Mount("/metrics", hh.Routes())

	ar := repositories.NewAccountRepository(db)
	tr := repositories.NewTokenRepository(db)
	auc := usecases.NewAccountUsecase(ar, tr)
	ah := handlers.NewAccountHandler(db, auc)
	r.Mount("/account", ah.Routes())

	http.ListenAndServe(":3000", r)

}
