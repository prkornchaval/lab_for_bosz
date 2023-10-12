package port

import (
	"context"
	"labForBosz/internal/core/domain"
)

type Repository interface {
	CreateCustomer(ctx context.Context, in domain.Customer) (*int, error)
	GetCustomer(ctx context.Context, id int) (*domain.Customer, error)
}
