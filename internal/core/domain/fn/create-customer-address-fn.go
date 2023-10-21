package fn

import (
	"context"
	"labForBosz/internal/core/domain"
)

type CreateCustomerAddressFn func(ctx context.Context, in domain.CreateCustomerAddress) (*int, error)
