package repository

import (
	"context"
	"errors"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/port"
)

var (
	customers = []domain.Customer{
		{Id: 1, FirstName: "user01_f", LastName: "user01_l", Age: 20},
	}
	countId = 1
)

type repository struct{}

func New() port.Repository {
	return &repository{}
}

func (r *repository) CreateCustomer(ctx context.Context, in domain.Customer) (*int, error) {
	countId++
	in.Id = countId
	customers = append(customers, in)
	return &in.Id, nil
}

func (r *repository) GetCustomer(ctx context.Context, id int) (*domain.Customer, error) {
	for _, customer := range customers {
		if customer.Id == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
