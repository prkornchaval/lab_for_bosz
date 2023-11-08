package cusrepo

import (
	"context"
	"database/sql"
	"errors"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/domain/fn"
	"labForBosz/internal/core/port"
)

type repo struct {
	dbPool *sql.DB
}

type TxnObjKey struct{}

func New(dbPool *sql.DB) port.CustRepository {
	return &repo{
		dbPool: dbPool,
	}
}

func (r *repo) CreateCustomerTx(ctx context.Context, in domain.Customer) (*int, error) {
	tx, ok := ctx.Value(TxnObjKey{}).(*sql.Tx)
	if !ok || tx == nil {
		return nil, errors.New("unable to get transactional from context")
	}

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
	if err := tx.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *repo) CreateAddressTx(ctx context.Context, in domain.Address) error {
	tx, ok := ctx.Value(TxnObjKey{}).(*sql.Tx)
	if !ok || tx == nil {
		return errors.New("unable to get transactional from context")
	}

	query := `insert into address
	(owner, subdistrict, district, province)
	values 
	($1, $2, $3, $4)`

	args := []interface{}{in.Owner, in.Subdistrict, in.District, in.Province}
	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (r *repo) CreateCustomerAddressTransactional(ctx context.Context, in domain.CreateCustomerAddress, f fn.CreateCustomerAddressFn) (*int, error) {
	tx, err := r.dbPool.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	txCtx := context.WithValue(ctx, TxnObjKey{}, tx)
	defer tx.Commit()

	id, err := f(txCtx, in)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return id, nil
}
