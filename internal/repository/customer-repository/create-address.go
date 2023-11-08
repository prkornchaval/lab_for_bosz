package customerrepo

import (
	"context"
	"labForBosz/internal/core/domain"
)

func (r *repository) CreateAddress(ctx context.Context, in domain.Address) error {
	query := `insert into address
	(owner, subdistrict, district, province)
	values 
	($1, $2, $3, $4)`

	args := []interface{}{in.Owner, in.Subdistrict, in.District, in.Province}
	if _, err := r.db.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}
