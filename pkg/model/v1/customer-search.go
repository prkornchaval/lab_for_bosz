package model

import "time"

type SearchCustomerRequest struct {
	Id int `json:"id" uri:"id"`
}

type CustomerRow struct {
	Id           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	MobileNo     string    `json:"mobile_no"`
	The1MemberId *string   `json:"the1_member_id"`
	The1MobileNo *string   `json:"the1_mobile_no"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    string    `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    string    `json:"updated_by"`
}
