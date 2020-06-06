package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/BranDebs/Avocado-Backend/account"
	"github.com/BranDebs/Avocado-Backend/api"
	"github.com/BranDebs/Avocado-Backend/repository/postgres"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	accSvc := setupAccountService()
	r := setupRouter()
	initRoutes(r, accSvc)
	runRouter(r)
}

func setupAccountService() account.AccountService {
	accPgSettings := postgres.ConnSettings{
		Host:     "avocadoro-db",
		Port:     5432,
		DBName:   "avocadoro",
		User:     "postgres",
		Password: "postgres123",
	}
	accRepo, err := postgres.NewRepository(accPgSettings)
	if err != nil {
		fmt.Printf("Error connecting to DB. Cause: %s", err)
		return nil
	}
	return account.NewAccountService(accRepo)
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r
}

func initRoutes(r *chi.Mux, svc account.AccountService) {
	h := api.NewHandler(svc)

	r.Get("/ping", h.Ping)

	r.Route("/accounts", func(r chi.Router) {
		r.Get("/", h.GetAccount)
		r.Post("/", h.PostAccount)
		r.Delete("/", h.DeleteAccount)
	})
}

func runRouter(r *chi.Mux) {
	errs := make(chan error, 2)

	go func() {
		log.Println("Running account service.")
		errs <- http.ListenAndServe(":8080", r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated (%s)", <-errs)
}
