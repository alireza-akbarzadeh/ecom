// Package products contains the product handlers for the ecom app.
package products

import (
	"log"
	"net/http"

	"github.com/techies/ecom/internal/json"
)

type handler struct {
	service Services
}

func NewHandler(service Services) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println("error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products := []string{"hello", "world"}

	json.Write(w, http.StatusOK, products)
}
