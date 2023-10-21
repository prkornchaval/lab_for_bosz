package model

type CreateCustomerRequest struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	MobileNo     string  `json:"mobile_no"`
	The1MemberId *string `json:"the1_member_id"`
	The1MobileNo *string `json:"the1_mobile_no"`
	CreatedBy    string  `json:"created_by"`
	UpdatedBy    string  `json:"updated_by"`
}

type CreateCustomerAddressRequest struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	MobileNo     string  `json:"mobile_no"`
	The1MemberId *string `json:"the1_member_id"`
	The1MobileNo *string `json:"the1_mobile_no"`
	Address      Address `json:"address"`
	CreatedBy    string  `json:"created_by"`
	UpdatedBy    string  `json:"updated_by"`
}

type Address struct {
	Subdistrict string `json:"subdistrict"`
	District    string `json:"district"`
	Province    string `json:"province"`
}

type CreateCustomerResponse struct {
	Id int `json:"id"`
}
