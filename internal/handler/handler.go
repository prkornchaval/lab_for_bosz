package handler

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/port"
	"labForBosz/pkg/model/v1"
)

type Handle interface {
	CreateCustomer(ctx context.Context, in model.CreateCustomerRequest) (*int, error)
	GetCustomer(ctx context.Context, id int) (*model.GetCustomerResponse, error)
	GetCustomerName(ctx context.Context, id int) (*model.GetCustomerNameResponse, error)
}

type handler struct {
	svc port.Service
}

func New(svc port.Service) Handle {
	return &handler{
		svc: svc,
	}
}

func (h *handler) CreateCustomer(ctx context.Context, in model.CreateCustomerRequest) (*int, error) {
	return h.svc.CreateCustomer(ctx, domain.Customer{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Age:       in.Age,
	})
}

func (h *handler) GetCustomer(ctx context.Context, id int) (*model.GetCustomerResponse, error) {
	res, err := h.svc.FindCustomer(ctx, id)
	if err != nil {
		return nil, err
	}
	// resp := model.GetCustomerResponse(*res)
	return &model.GetCustomerResponse{
		Id:        res.Id,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Age:       res.Age,
	}, nil
}

func (h *handler) GetCustomerName(ctx context.Context, id int) (*model.GetCustomerNameResponse, error) {
	res, err := h.svc.FindCustomer(ctx, id)
	if err != nil {
		return nil, err
	}
	// resp :=  model.GetCustomerNameResponse(*res)
	return &model.GetCustomerNameResponse{
		Id:        res.Id,
		FirstName: res.FirstName,
	}, nil
}
