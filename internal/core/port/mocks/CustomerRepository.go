// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "labForBosz/internal/core/domain"
	fn "labForBosz/internal/core/domain/fn"

	mock "github.com/stretchr/testify/mock"
)

// CustomerRepository is an autogenerated mock type for the CustomerRepository type
type CustomerRepository struct {
	mock.Mock
}

// CreateAddress provides a mock function with given fields: ctx, in
func (_m *CustomerRepository) CreateAddress(ctx context.Context, in domain.Address) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Address) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCustomer provides a mock function with given fields: ctx, in
func (_m *CustomerRepository) CreateCustomer(ctx context.Context, in domain.Customer) (*int, error) {
	ret := _m.Called(ctx, in)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Customer) (*int, error)); ok {
		return rf(ctx, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Customer) *int); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Customer) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCustomerAddressTransactional provides a mock function with given fields: ctx, in, f
func (_m *CustomerRepository) CreateCustomerAddressTransactional(ctx context.Context, in domain.CreateCustomerAddress, f fn.CreateCustomerAddressFn) (*int, error) {
	ret := _m.Called(ctx, in, f)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateCustomerAddress, fn.CreateCustomerAddressFn) (*int, error)); ok {
		return rf(ctx, in, f)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateCustomerAddress, fn.CreateCustomerAddressFn) *int); ok {
		r0 = rf(ctx, in, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CreateCustomerAddress, fn.CreateCustomerAddressFn) error); ok {
		r1 = rf(ctx, in, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomer provides a mock function with given fields: ctx, id
func (_m *CustomerRepository) GetCustomer(ctx context.Context, id *int) ([]domain.Customer, error) {
	ret := _m.Called(ctx, id)

	var r0 []domain.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *int) ([]domain.Customer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *int) []domain.Customer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCustomer provides a mock function with given fields: ctx, in
func (_m *CustomerRepository) UpdateCustomer(ctx context.Context, in domain.Customer) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Customer) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCustomerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCustomerRepository creates a new instance of CustomerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCustomerRepository(t mockConstructorTestingTNewCustomerRepository) *CustomerRepository {
	mock := &CustomerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}