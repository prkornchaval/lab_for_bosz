package customerrepo

import (
	"context"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/port"
	"labForBosz/pkg/dbctx"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type repository struct {
	db      *dbctx.DB
	scanapi *pgxscan.API
}

func New(db *dbctx.DB, scanapi *pgxscan.API) port.CustomerRepository {
	return &repository{
		db:      db,
		scanapi: scanapi,
	}
}

func (r *repository) GetCustomer(ctx context.Context, id *int) ([]domain.Customer, error) {
	queryString := `select * from customer where id = $1`
	rows, err := r.db.Query(ctx, queryString, id)
	if err != nil {
		return nil, err
	}

	var resp []domain.Customer
	if err := r.scanapi.ScanAll(&resp, rows); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *repository) UpdateCustomer(ctx context.Context, in domain.Customer) error {
	query := `update customer set
			firstname = $2,
			lastname = $3,
			mobile_no = $4,
			the1_member_id = $5,
			the1_mobile_no = $6,
			updated_at = $7,
			updated_by = $8
		where id = $1`
	args := []interface{}{
		in.Id,
		in.FirstName,
		in.LastName,
		in.MobileNo,
		in.The1MemberId,
		in.The1MobileNo,
		in.UpdatedAt,
		in.UpdatedBy,
	}
	_, err := r.db.Exec(ctx, query, args...)
	return err
}
