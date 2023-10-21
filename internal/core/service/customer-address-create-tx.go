package service

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/domain/fn"
	"labForBosz/pkg/dbctx"
)

func (s *service) CreateCustomerAddressTransactional(ctx context.Context, in domain.CreateCustomerAddress, f fn.CreateCustomerAddressFn) (*int, error) {
	var id int
	err := dbctx.Transactional(ctx, func(txCtx context.Context) error {
		res, err := f(txCtx, in)
		if err != nil {
			return err
		}
		id = *res
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &id, nil
}
