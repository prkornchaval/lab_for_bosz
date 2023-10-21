package service

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/port"
)

type service struct {
	customerRepo port.CustomerRepository
}

func New(customerRepo port.CustomerRepository) port.CustomerService {
	return &service{
		customerRepo: customerRepo,
	}
}

func (s *service) GetCustomer(ctx context.Context) ([]domain.Customer, error) {
	return s.customerRepo.GetCustomer(ctx, nil)
}

func (s *service) GetCustomerbyId(ctx context.Context, id int) (*domain.Customer, error) {
	res, err := s.customerRepo.GetCustomer(ctx, &id)
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (s *service) UpdateCustomer(ctx context.Context, in domain.Customer) error {
	return s.customerRepo.UpdateCustomer(ctx, in)
}
