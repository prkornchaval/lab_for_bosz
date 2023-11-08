package customerrepo

import (
	"context"
	"labForBosz/internal/core/domain"
)

func (r *repository) CreateCustomer(ctx context.Context, in domain.Customer) (*int, error) {
	query := `insert into customer (
		firstname, lastname, mobile_no,
		the1_member_id, the1_mobile_no, 
		created_at, created_by,
		updated_at, updated_by
	) values (
		$1, $2,
		$3, $4,
		$5, $6, 
		$7, $8,
		$9
	) RETURNING id`

	args := []interface{}{
		in.FirstName, in.LastName, in.MobileNo,
		in.The1MemberId, in.The1MobileNo,
		in.CreatedAt, in.CreatedBy,
		in.UpdatedAt, in.UpdatedBy,
	}

	var id int
	if err := r.db.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return nil, err
	}
	return &id, nil
}
