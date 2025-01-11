package main

import (
	"context"
	"net/http"

	"github.com/0ne290/JavaCodeTestTask/internal/core/domain"
	"github.com/0ne290/JavaCodeTestTask/internal/infrastructure"
	walletHandlers "github.com/0ne290/JavaCodeTestTask/internal/web/handlers/wallet"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	uuidProviderFactory = func() domain.UuidProvider {
		return &infrastructure.UuidProvider{}
	}

	unitOfWorkFactory func() domain.UnitOfWork
)

func main() {
	pool, err := pgxpool.New(context.Background(), "user=root password=dcba host=localhost port=5432 dbname=java_code_test_task")

	if err != nil {
		panic(err.Error())
	}

	unitOfWorkFactory = func() domain.UnitOfWork {
		return infrastructure.NewUnitOfWork(context.Background(), pool)
	}

	r := chi.NewRouter()

	r.Mount("/api", Versions())

	srv := &http.Server{
		Addr:    "192.168.0.101:80",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}

func Versions() chi.Router {
	r := chi.NewRouter()

	r.Mount("/v1", V1())

	return r
}

func V1() chi.Router {
	r := chi.NewRouter()

	r.Route("/wallets", func(r chi.Router) {
		r.Post("/", walletHandlers.Create(uuidProviderFactory, unitOfWorkFactory))
	})

	return r
}
