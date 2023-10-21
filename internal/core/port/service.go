package port

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/domain/fn"
)

type CustomerService interface {
	GetCustomer(ctx context.Context) ([]domain.Customer, error)
	GetCustomerbyId(ctx context.Context, id int) (*domain.Customer, error)
	UpdateCustomer(ctx context.Context, in domain.Customer) error

	CreateCustomer(ctx context.Context, in domain.Customer) (*int, error)
	CreateCustomerAddress(ctx context.Context, in domain.CreateCustomerAddress) (*int, error)
	CreateCustomerAddressTransactional(ctx context.Context, in domain.CreateCustomerAddress, f fn.CreateCustomerAddressFn) (*int, error)
}
