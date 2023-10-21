package handler

import (
	"context"
	"errors"
	"labForBosz/pkg/model/v1"
)

// @Summary      get customer by id
// @Tags         customer
// @Accept       json
// @Security 	 ApiKeyAuth
// @Param 		 id path string false "id"
// @Produce      json
// @Success      200 {object} model.CustomerRow
// @Router		 /v1/{id} [get]
func (h *handler) GetCustomerbyId(ctx context.Context, in model.SearchCustomerRequest) (*model.CustomerRow, error) {
	res, err := h.cusSvc.GetCustomerbyId(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("not found")
	}
	resp := model.CustomerRow(*res)
	return &resp, nil
}
