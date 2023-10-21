package dbctx

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	queryKey struct{}
	dbKey    struct{}
)

// create query type interface
type query interface {
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
}

func NewContext(parent context.Context, db *pgxpool.Pool) context.Context {
	ctx := context.WithValue(parent, queryKey{}, db)
	ctx = context.WithValue(ctx, dbKey{}, db)
	return ctx
}

func GinMiddleware(db *pgxpool.Pool) func(*gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = NewContext(ctx, db)
		c.Request = c.Request.WithContext(ctx)
	}
}

func HttpMiddleware(db *pgxpool.Pool) func(handler http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = NewContext(ctx, db)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

// getQuery: return query interface
func getQuery(ctx context.Context) query {
	dbPool, ok := ctx.Value(queryKey{}).(*pgxpool.Pool)
	if ok {
		return dbPool
	}
	return ctx.Value(queryKey{}).(*pgxpool.Tx)
}

// getDB: return dbPool
func getDB(ctx context.Context) *pgxpool.Pool {
	return ctx.Value(dbKey{}).(*pgxpool.Pool)
}

func QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return getQuery(ctx).QueryRow(ctx, query, args...)
}

func Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return getQuery(ctx).Query(ctx, query, args...)
}

func Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return getQuery(ctx).Exec(ctx, query, args...)
}

func Transactional(ctx context.Context, f func(context.Context) error) error {
	db := getDB(ctx)
	tx, err := db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return errors.New("unable to get transactional from context")
	}
	txctx := context.WithValue(ctx, queryKey{}, tx)
	defer tx.Rollback(txctx)
	if err := f(txctx); err != nil {
		return err
	}
	return tx.Commit(txctx)
}
