package products

import (
	"context"

	repo "github.com/techies/ecom/internal/adapters/postgres/sqlc"
)

type Services interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	FindProductByID(ctx context.Context, id int64) (repo.Product, error)
	CreateProduct(ctx context.Context, product repo.CreateProductParams) (repo.Product, error)
	UpdateProduct(ctx context.Context, product repo.UpdateProductParams, id int64) (repo.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Services {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)

}

func (s *svc) FindProductByID(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductByID(ctx, id)
}

func (s *svc) CreateProduct(ctx context.Context, product repo.CreateProductParams) (repo.Product, error) {
	return s.repo.CreateProduct(ctx, product)
}

func (s *svc) UpdateProduct(ctx context.Context, product repo.UpdateProductParams, id int64) (repo.Product, error) {
	return s.repo.UpdateProduct(ctx, repo.UpdateProductParams{
		Name:         product.Name,
		PriceInCents: product.PriceInCents,
		Quantity:     product.Quantity,
		ID:           id,
	})
}

func (s *svc) DeleteProduct(ctx context.Context, id int64) error {
	return s.repo.DeleteProduct(ctx, id)
}
