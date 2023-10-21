package handler

import (
	"context"
	"labForBosz/internal/core/port"
	"labForBosz/pkg/model/v1"
)

type Handler interface {
	CreateCustomer(ctx context.Context, in model.CreateCustomerRequest) (*model.CreateCustomerResponse, error)
	CreateCustomerAddress(ctx context.Context, in model.CreateCustomerAddressRequest) (*model.CreateCustomerResponse, error)
	GetCustomerbyId(ctx context.Context, in model.SearchCustomerRequest) (*model.CustomerRow, error)
}

type handler struct {
	cusSvc port.CustomerService
}

func New(cusSvc port.CustomerService) Handler {
	return &handler{
		cusSvc: cusSvc,
	}
}
