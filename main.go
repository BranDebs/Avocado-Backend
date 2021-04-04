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
	"github.com/BranDebs/Avocado-Backend/task"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	config, err := newConfig()
	if err != nil {
		log.Fatalf("Unable to initialise config: %s", err)
	}

	accSvc, taskSvc, err := setupServices(config)
	if err != nil {
		log.Fatalf("Unable to setup account service: %s", err)
	}

	router := setupRouter()
	initRoutes(router, accSvc, taskSvc)
	runRouter(router, config)
}

func setupServices(c configer) (account.Service, task.Service, error) {
	var dbSettings postgres.ConnSettings

	if err := c.unmarshalKey("db", &dbSettings); err != nil {
		return nil, nil, fmt.Errorf("setup services: load in db config: %w", err)
	}

	accRepo, err := postgres.NewAccountRepository(dbSettings)
	if err != nil {
		return nil, nil, fmt.Errorf("setup services: create account repository: %w", err)
	}

	var jwtSettings account.JWTSettings

	if err := c.unmarshalKey("jwt", &jwtSettings); err != nil {
		return nil, nil, fmt.Errorf("setup services: load in jwt config: %w", err)
	}

	taskRepo, err := postgres.NewTaskRepository(dbSettings)
	if err != nil {
		return nil, nil, fmt.Errorf("setup services: create task repository: %w", err)
	}

	return account.NewService(accRepo, &jwtSettings), task.NewService(taskRepo), nil
}

func setupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	return router
}

func initRoutes(router *chi.Mux, account account.Service, task task.Service) {
	handler := api.NewHandler(account, task)

	router.Get("/ping", handler.Ping)

	router.Route("/accounts", func(r chi.Router) {
		r.Post("/login", handler.LoginAccount)
		r.Post("/", handler.CreateAccount)
		r.Delete("/{email}", handler.DeleteAccount)
	})

	router.Route("/tasks", func(r chi.Router) {
		r.Post("/", handler.CreateTask)
		r.Get("/", handler.FindTasks)
		r.Put("/{id}", handler.UpdateTask)
		r.Delete("/{id}", handler.DeleteTask)
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
