package customerrepo

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/domain/fn"
)

func (r *repository) CreateCustomerAddressTransactional(ctx context.Context, in domain.CreateCustomerAddress, f fn.CreateCustomerAddressFn) (*int, error) {
	var id int
	err := r.db.Transactional(ctx, func(txCtx context.Context) error {
		res, err := f(txCtx, in)
		// TODO MOCK ERROR
		//err = errors.New("text")
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
