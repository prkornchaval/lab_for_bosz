package dbctx

import (
	"context"
	"errors"
	"fmt"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	queryKey struct{}
	dbKey    struct{}
	ct       context.Context
}

func OpenPgxPool(ctx context.Context, dns string) (db *DB, err error) {
	cfg, err := pgxpool.ParseConfig(dns)
	if err != nil {
		return nil, err
	}
	cfg.ConnConfig.Tracer = otelpgx.NewTracer()
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	// ctx = NewContext(ctx, pool)
	var d DB
	d.ct = context.WithValue(ctx, d.queryKey, pool)
	d.ct = context.WithValue(d.ct, d.dbKey, pool)
	return &d, nil
}

func (d *DB) getQuery(ctx context.Context) query {
	if tx, ok := ctx.Value(d.dbKey).(*pgxpool.Tx); ok {
		return tx
	}

	dbPool, ok := d.ct.Value(d.queryKey).(*pgxpool.Pool)
	if ok {
		return dbPool
	}
	return d.ct.Value(d.queryKey).(*pgxpool.Tx)
}

func (d *DB) getDB(ctx context.Context) *pgxpool.Pool {
	return d.ct.Value(d.dbKey).(*pgxpool.Pool)
}

func (d *DB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.getQuery(ctx).QueryRow(ctx, query, args...)
}

func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.getQuery(ctx).Query(ctx, query, args...)
}

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.getQuery(ctx).Exec(ctx, query, args...)
}

func (d *DB) Transactional(ctx context.Context, f func(context.Context) error) error {
	db := d.getDB(ctx)
	tx, err := db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return errors.New("unable to get transactional from context")
	}
	fmt.Println("1", d.ct.Value(d.queryKey))
	txctx := context.WithValue(d.ct, d.queryKey, tx)
	fmt.Println("2", d.ct.Value(d.queryKey))
	defer tx.Rollback(txctx)
	if err := f(txctx); err != nil {
		return err
	}
	return tx.Commit(txctx)
}
