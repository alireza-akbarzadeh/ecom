package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/techies/ecom/internal/products"
)

type application struct {
	config config
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

// mount entry point of the ecom app
func (app *application) mount() http.Handler {
	// init chi router
	r := chi.NewRouter()

	// middlewares

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// routes

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good! i guess"))
	})
	productService := products.NewService()
	productHandler := products.NewHandler(productService)

	r.Get("/products", productHandler.ListProductHandler)

	return r
}

// run responsible for running the application
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at addr %s", app.config.addr)
	return srv.ListenAndServe()

}
