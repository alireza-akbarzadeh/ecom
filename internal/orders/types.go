package orders

import (
	"context"

	repo "github.com/techies/ecom/internal/adapters/postgres/sqlc"
)

type OrderItem struct {
	ProductID int64 `json:"productId"`
	Quantity  int32 `json:"quantity"`
}

type CreateOrderParams struct {
	CustomerID int64       `json:"customerId"`
	Items      []OrderItem `json:"items"`
}

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder CreateOrderParams) (repo.Order, error)
}
