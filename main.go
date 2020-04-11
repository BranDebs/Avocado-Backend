package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/BranDebs/Avocado-Backend/api"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := setupRouter()
	initRoutes(r)
	runRouter(r)
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r
}

func initRoutes(r *chi.Mux) {
	h := api.NewHandler(nil)

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
