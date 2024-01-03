package cusservice

import (
	"context"
	"errors"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/port"
	"labForBosz/internal/core/port/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testingModule struct {
	repo *mocks.CustRepository
	svc  port.CustService
	util *mockServiceUtils
}

func newTestingModule(t *testing.T) testingModule {
	repo := mocks.NewCustRepository(t)
	util := newMockServiceUtils(t)
	return testingModule{
		repo: repo,
		util: util,
		svc:  New(repo, util),
	}
}

func TestCreateCustomerAddressFn(t *testing.T) {
	ti := time.Now()
	var (
		input = domain.CreateCustomerAddress{
			Customer: domain.Customer{
				FirstName:    "test_tx_name",
				LastName:     "test_tx_last-name",
				MobileNo:     "test_tx_mobile-no",
				The1MemberId: nil,
				The1MobileNo: nil,
				IsActive:     true,
				CreatedAt:    ti,
				CreatedBy:    "me",
				UpdatedAt:    ti,
				UpdatedBy:    "me",
			},
			Address: domain.Address{
				Subdistrict: "test_sub_district",
				District:    "test_district",
				Province:    "test_province",
			},
		}
	)
	t.Run("00_util_fails", func(t *testing.T) {
		var (
			ctx = context.Background()
			tm  = newTestingModule(t)
		)

		tm.util.On("createCustomers", tm.svc).Return(errors.New("util failed"))
		_, err := tm.svc.CreateCustomerAddressFn(ctx, input)

		assert.Error(t, err)
	})
	t.Run("01_create_customer_error", func(t *testing.T) {
		var (
			ctx   = context.Background()
			tm    = newTestingModule(t)
			ti    = time.Now()
			input = domain.CreateCustomerAddress{
				Customer: domain.Customer{
					FirstName:    "test_tx_name",
					LastName:     "test_tx_last-name",
					MobileNo:     "test_tx_mobile-no",
					The1MemberId: nil,
					The1MobileNo: nil,
					IsActive:     true,
					CreatedAt:    ti,
					CreatedBy:    "me",
					UpdatedAt:    ti,
					UpdatedBy:    "me",
				},
				Address: domain.Address{
					Subdistrict: "test_sub_district",
					District:    "test_district",
					Province:    "test_province",
				},
			}
			dbResponseId = 10
			_            = dbResponseId
		)

		tm.util.On("createCustomers").Return(nil)
		tm.repo.On("CreateCustomerTx", ctx, input.Customer).Return(nil, errors.New("test"))

		_, err := tm.svc.CreateCustomerAddressFn(ctx, input)

		assert.Error(t, err)
	})

	t.Run("02_create_customer_error_id_is_null", func(t *testing.T) {
		var (
			ctx = context.Background()
			tm  = newTestingModule(t)
		)

		tm.util.On("createCustomers").Return(nil)
		tm.repo.On("CreateCustomerTx", ctx, input.Customer).Return(nil, nil)

		_, err := tm.svc.CreateCustomerAddressFn(ctx, input)

		assert.Error(t, err)
	})

	t.Run("03__create_address_error", func(t *testing.T) {
		var (
			ctx = context.Background()
			tm  = newTestingModule(t)
			id  = 20
		)

		tm.util.On("createCustomers").Return(nil)
		tm.repo.On("CreateCustomerTx", ctx, input.Customer).Return(&id, nil)
		input.Address.Owner = id
		tm.repo.On("CreateAddressTx", ctx, input.Address).Return(errors.New("create address error"))

		_, err := tm.svc.CreateCustomerAddressFn(ctx, input)
		assert.EqualError(t, errors.New("create address error"), err.Error())
	})

	t.Run("04_success", func(t *testing.T) {
		var (
			ctx = context.Background()
			tm  = newTestingModule(t)
			id  = 20
		)

		tm.util.On("createCustomers").Return(nil)
		tm.repo.On("CreateCustomerTx", ctx, input.Customer).Return(&id, nil)
		input.Address.Owner = id
		tm.repo.On("CreateAddressTx", ctx, input.Address).Return(nil)

		_, err := tm.svc.CreateCustomerAddressFn(ctx, input)
		assert.NoError(t, err)
	})
}

func TestCreateCustomerAddressTransaction(t *testing.T) {
	ti := time.Now()
	var (
		input = domain.CreateCustomerAddress{
			Customer: domain.Customer{
				FirstName:    "test_tx_name",
				LastName:     "test_tx_last-name",
				MobileNo:     "test_tx_mobile-no",
				The1MemberId: nil,
				The1MobileNo: nil,
				IsActive:     true,
				CreatedAt:    ti,
				CreatedBy:    "me",
				UpdatedAt:    ti,
				UpdatedBy:    "me",
			},
			Address: domain.Address{
				Subdistrict: "test_sub_district",
				District:    "test_district",
				Province:    "test_province",
			},
		}
	)
	t.Run("01_fail", func(t *testing.T) {
		var (
			ctx = context.Background()
			tm  = newTestingModule(t)
		)

		tm.repo.On("CreateCustomerAddressTransactional", ctx, input, mock.AnythingOfType("fn.CreateCustomerAddressFn")).
			Return(nil, errors.New("mock error"))
		id, err := tm.svc.CreateCustomerAddressTransaction(ctx, input)
		assert.Error(t, err)
		assert.Nil(t, id)
	})

	t.Run("02_success", func(t *testing.T) {
		var (
			ctx   = context.Background()
			tm    = newTestingModule(t)
			cusId = 20
		)

		tm.repo.On("CreateCustomerAddressTransactional", ctx, input, mock.AnythingOfType("fn.CreateCustomerAddressFn")).
			Return(&cusId, nil)
		id, err := tm.svc.CreateCustomerAddressTransaction(ctx, input)
		assert.NoError(t, err)
		assert.Equal(t, 20, *id)
	})
}
