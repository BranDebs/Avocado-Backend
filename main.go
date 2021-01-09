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
	config, err := newConfig()
	if err != nil {
		log.Fatalf("Unable to initialise config: %s", err)
	}

	accSvc, err := setupAccountService(config)
	if err != nil {
		log.Fatalf("Unable to setup account service: %s", err)
	}

	router := setupRouter()
	initRoutes(router, accSvc)
	runRouter(router, config)
}

func setupAccountService(c configer) (account.AccountService, error) {
	var accSettings postgres.ConnSettings

	if err := c.unmarshalKey("db", &accSettings); err != nil {
		return nil, fmt.Errorf("setup account service: load in config values: %w", err)
	}

	accRepo, err := postgres.NewRepository(accSettings)
	if err != nil {
		return nil, fmt.Errorf("setup account service: create account repository: %w", err)
	}

	jwtTTL := c.getInt64("jwt.ttl")

	return account.NewAccountService(accRepo, jwtTTL), nil
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
		r.Post("/login", handler.LoginAccount)
		r.Post("/", handler.CreateAccount)
		r.Delete("/{email}", handler.DeleteAccount)
	})
}

func runRouter(r *chi.Mux, c configer) {
	errs := make(chan error, 2)

	go func() {
		listenAddr := c.getString("app.listening_addr")
		log.Printf("Running account service on (%s).\n", listenAddr)
		errs <- http.ListenAndServe(listenAddr, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated (%s)", <-errs)
}
