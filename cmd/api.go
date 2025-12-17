package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/techies/ecom/internal/adapters/postgres/sqlc"
	"github.com/techies/ecom/internal/orders"
	"github.com/techies/ecom/internal/products"
)

type application struct {
	config config
	db     *pgx.Conn
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
	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)

	// product routes
	r.Get("/products", productHandler.ListProductHandler)
	r.Get("/products/:id", productHandler.GetProductHandler)
	r.Post("/products", productHandler.CreateProductHandler)
	r.Put("/products/:id", productHandler.UpdateProductHandler)
	r.Delete("/products/:id", productHandler.DeleteProductHandler)

	orderService := orders.NewService(repo.New(app.db), app.db)
	orderHandler := orders.NewHandler(orderService)
	// order routes
	r.Post("/order", orderHandler.PlaceOrder)

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
