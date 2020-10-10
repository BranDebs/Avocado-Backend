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
	router := setupRouter()
	initRoutes(router, accSvc)
	runRouter(router)
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
		log.Printf("Error connecting to DB: %s", err)
		return nil
	}
	return account.NewAccountService(accRepo)
}

func setupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	return router
}

func initRoutes(router *chi.Mux, svc account.AccountService) {
	handler := api.NewHandler(svc)

	router.Get("/ping", handler.Ping)

	router.Route("/accounts", func(r chi.Router) {
		r.Get("/", handler.GetAccount)
		r.Post("/", handler.PostAccount)
		r.Delete("/", handler.DeleteAccount)
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
