package handler

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/pkg/model/v1"
	"time"
)

// @Summary      create customer and address
// @Tags         customer
// @Accept       json
// @Security 	 ApiKeyAuth
// @Param 		 create-customer-request body model.CreateCustomerAddressRequest false "Create customer"
// @Produce      json
// @Success      200 {object} model.CreateCustomerResponse
// @Router		 /v1/customer-address [post]
func (h *handler) CreateCustomerAddress(ctx context.Context, in model.CreateCustomerAddressRequest) (*model.CreateCustomerResponse, error) {
	id, err := h.cusSvc.CreateCustomerAddressTransactional(ctx, domain.CreateCustomerAddress{
		Customer: domain.Customer{
			FirstName:    in.FirstName,
			LastName:     in.LastName,
			MobileNo:     in.MobileNo,
			The1MemberId: in.The1MemberId,
			The1MobileNo: in.The1MobileNo,
			CreatedAt:    time.Now(),
			CreatedBy:    in.CreatedBy,
			UpdatedAt:    time.Now(),
			UpdatedBy:    in.UpdatedBy,
		},
		Address: domain.Address{
			Subdistrict: in.Address.Subdistrict,
			District:    in.Address.District,
			Province:    in.Address.Province,
		},
	}, h.cusSvc.CreateCustomerAddress)

	if err != nil {
		return nil, err
	}

	return &model.CreateCustomerResponse{Id: *id}, nil

}
