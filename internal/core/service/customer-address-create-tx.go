package service

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/domain/fn"
)

func (s *service) CreateCustomerAddressTransactional(ctx context.Context, in domain.CreateCustomerAddress, f fn.CreateCustomerAddressFn) (*int, error) {
	id, err := s.customerRepo.CreateCustomerAddressTransactional(ctx, in, s.CreateCustomerAddress)
	if err != nil {
		return nil, err
	}
	return id, nil
}
