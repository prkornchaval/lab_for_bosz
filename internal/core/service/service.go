package service

import (
	"context"
	"database/sql"
	"errors"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/port"
)

type service struct {
	repo port.Repository
}

func New(repo port.Repository) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateCustomer(ctx context.Context, in domain.Customer) (*int, error) {
	if in.FirstName == "" {
		return nil, errors.New("first name require")
	}
	return s.repo.CreateCustomer(ctx, in)
}

func (s *service) FindCustomer(ctx context.Context, id int) (*domain.Customer, error) {
	res, err := s.repo.GetCustomer(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("data not found")
		}
		return nil, err
	}
	if res == nil {
		return nil, errors.New("data not found")
	}
	return res, nil
}
