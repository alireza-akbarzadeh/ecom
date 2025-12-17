// Package products contains the product handlers for the ecom app.
package products

import (
	"encoding/json"
	"net/http"

	repo "github.com/techies/ecom/internal/adapters/postgres/sqlc"
	"github.com/techies/ecom/internal/utils"
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
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		utils.WriteJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.WriteToJSON(w, http.StatusOK, products)
}

func (h *handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r)
	if err != nil {
		utils.WriteJSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.FindProductByID(r.Context(), id)
	if err != nil {
		utils.WriteJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.WriteToJSON(w, http.StatusOK, product)
}

func (h *handler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var req repo.CreateProductParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateProduct(r.Context(), req)
	if err != nil {
		utils.WriteJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.WriteToJSON(w, http.StatusCreated, product)
}

func (h *handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r)
	if err != nil {
		utils.WriteJSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var productParams repo.UpdateProductParams
	err = json.NewDecoder(r.Body).Decode(&productParams)
	if err != nil {
		utils.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updatedProduct, err := h.service.UpdateProduct(r.Context(), productParams, id)
	if err != nil {
		utils.WriteJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	utils.WriteToJSON(w, http.StatusOK, updatedProduct)
}

func (h *handler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r)
	if err != nil {
		utils.WriteJSONError(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = h.service.DeleteProduct(r.Context(), id)
	if err != nil {
		utils.WriteJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
