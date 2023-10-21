package service

import (
	"context"
	"errors"
	"fmt"
	"labForBosz/internal/core/domain"
)

func (s service) CreateCustomerAddress(ctx context.Context, in domain.CreateCustomerAddress) (*int, error) {
	id, err := s.customerRepo.CreateCustomer(ctx, in.Customer)
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, errors.New("create custoemr no id")
	}

	fmt.Printf("\n\n>>%v<<\n\n", *id)

	if err := s.customerRepo.CreateAddress(ctx, domain.Address{
		Owner:       *id,
		Subdistrict: in.Address.Subdistrict,
		District:    in.Address.District,
		Province:    in.Address.Province,
	}); err != nil {
		return nil, err
	}
	//return nil, errors.New("mock error")
	return id, nil
}
