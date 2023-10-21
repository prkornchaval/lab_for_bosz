package mapperhandler

import (
	"labForBosz/internal/core/domain"
	"labForBosz/pkg/model/v1"
	"time"
)

func MapCreateCustomerRequestModel(mo model.CreateCustomerRequest) domain.Customer {
	return domain.Customer{
		FirstName:    mo.FirstName,
		LastName:     mo.LastName,
		MobileNo:     mo.MobileNo,
		The1MemberId: mo.The1MemberId,
		The1MobileNo: mo.The1MobileNo,
		CreatedAt:    time.Now(),
		CreatedBy:    mo.CreatedBy,
		UpdatedAt:    time.Now(),
		UpdatedBy:    mo.UpdatedBy,
	}
}
