package handler

import (
	"context"
	"errors"
	mapperhandler "labForBosz/internal/handler/mapper-handler"
	"labForBosz/pkg/model/v1"
)

// @Summary      create customer
// @Tags         customer
// @Accept       json
// @Security 	 ApiKeyAuth
// @Param 		 create-customer-request body model.CreateCustomerRequest false "Create customer"
// @Produce      json
// @Success      200 {object} model.CreateCustomerResponse
// @Router		 /v1 [post]
func (h *handler) CreateCustomer(ctx context.Context, in model.CreateCustomerRequest) (*model.CreateCustomerResponse, error) {
	id, err := h.cusSvc.CreateCustomer(ctx, mapperhandler.MapCreateCustomerRequestModel(in))
	if err != nil {
		return nil, err
	}
	if id == nil {
		return nil, errors.New("no id response")
	}
	return &model.CreateCustomerResponse{
		Id: *id,
	}, nil
}
