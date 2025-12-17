// Package utils provides utility functions for common tasks.
package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ParseID extracts and parses the "id" URL parameter from the request.
func ParseID(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.ParseInt(idStr, 10, 64)
}

// WriteToJSON writes the given data as a JSON response with the specified status code.
func WriteToJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteJSONError sends a JSON error response with the specified status code
func WriteJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

// Read decodes JSON from the request body into the provided destination structure.
func Read(r *http.Request, dst any) error {
	decoded := json.NewDecoder(r.Body)
	decoded.DisallowUnknownFields()
	return decoded.Decode(dst)
}
