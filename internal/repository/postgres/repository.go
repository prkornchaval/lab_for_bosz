package postgres

import (
	"context"
	"database/sql"
	"labForBosz/internal/core/domain"
	"labForBosz/internal/core/port"

	"github.com/georgysavva/scany/sqlscan"
)

type repository struct {
	db      *sql.DB
	scanApi *sqlscan.API
}

func New(db *sql.DB, scanApi *sqlscan.API) port.Repository {
	return &repository{
		db:      db,
		scanApi: scanApi,
	}
}

func (r *repository) CreateCustomer(ctx context.Context, in domain.Customer) (*int, error) { // returnin (id)
	sqlquery := `insert into users (firstname, lastname, age)  values($1, $2, $3)`
	args := []interface{}{in.FirstName, in.LastName, in.Age}
	result, err := r.db.ExecContext(ctx, sqlquery, args...)
	if err != nil {
		return nil, err
	}
	resId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	id := int(resId)
	return &id, nil
}

func (r *repository) GetCustomer(ctx context.Context, id int) (*domain.Customer, error) {
	sqlquery := `select id, firstname, lastname, age from users where $1`
	rows, err := r.db.QueryContext(ctx, sqlquery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customer domain.Customer
	err = r.scanApi.ScanAll(&customer, rows)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
