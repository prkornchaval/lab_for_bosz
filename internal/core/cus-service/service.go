package cusservice

import (
	"context"
	"errors"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/port"
)

type service struct {
	util serviceUtils
	repo port.CustRepository
}

type serviceUtilImpl struct{}

type serviceUtils interface {
	createCustomers(s *service) error
}

func NewServiceUtil() serviceUtilImpl {
	return serviceUtilImpl{}
}

func New(repo port.CustRepository, util serviceUtils) port.CustService {
	return &service{
		util: util,
		repo: repo,
	}
}

func (s *service) CreateCustomerAddressTransaction(ctx context.Context, in domain.CreateCustomerAddress) (*int, error) {
	return s.repo.CreateCustomerAddressTransactional(ctx, in, s.CreateCustomerAddressFn)
}

func (s *service) CreateCustomerAddressFn(ctx context.Context, in domain.CreateCustomerAddress) (*int, error) {
	err := s.util.createCustomers(s)
	if err != nil {
		return nil, err
	}

	id, err := s.repo.CreateCustomerTx(ctx, in.Customer)
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, errors.New("create custoemr no id")
	}

	// TODO force error for test rollback
	// return nil, errors.New("mock error")

	in.Address.Owner = *id
	if err := s.repo.CreateAddressTx(ctx, in.Address); err != nil {
		return nil, err
	}
	return id, nil
}
