package products

import "context"

type Services interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
}

func NewService() Services {
	return &svc{}
}

func (s *svc) ListProducts(ctx context.Context) error {

	return nil
}
