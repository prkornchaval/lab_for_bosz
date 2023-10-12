package port

import (
	"context"
	"labForBosz/internal/core/domain"
)

type Service interface {
	CreateCustomer(ctx context.Context, in domain.Customer) (*int, error)
	FindCustomer(ctx context.Context, id int) (*domain.Customer, error)
}
