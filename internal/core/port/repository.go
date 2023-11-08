package port

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/domain/fn"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, in domain.Customer) (*int, error)
	CreateAddress(ctx context.Context, in domain.Address) error
	GetCustomer(ctx context.Context, id *int) ([]domain.Customer, error)
	UpdateCustomer(ctx context.Context, in domain.Customer) error
	CreateCustomerAddressTransactional(ctx context.Context, in domain.CreateCustomerAddress, f fn.CreateCustomerAddressFn) (*int, error)
}
