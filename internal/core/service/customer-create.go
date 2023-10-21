package service

import (
	"context"
	"labForBosz/internal/core/domain"
)

func (s *service) CreateCustomer(ctx context.Context, in domain.Customer) (*int, error) {
	return s.customerRepo.CreateCustomer(ctx, in)
}
