package domain

import "time"

type Customer struct {
	Id           int       `db:"id"`
	FirstName    string    `db:"firstname"`
	LastName     string    `db:"lastname"`
	MobileNo     string    `db:"mobile_no"`
	The1MemberId *string   `db:"the1_member_id"`
	The1MobileNo *string   `db:"the1_mobile_no"`
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	CreatedBy    string    `db:"created_by"`
	UpdatedAt    time.Time `db:"updated_at"`
	UpdatedBy    string    `db:"updated_by"`
}
