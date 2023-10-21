package infrastructure

import (
	"context"
	"log"

	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

// connString example: "host=localhost port=5432 database=order user=postgres password=secret"
func NewPostgresWithScanApi(ctx context.Context, connString string) (*pgxpool.Pool, *pgxscan.API) {
	db := NewPostgresDB(ctx, connString)
	scanapi := NewScanApi()
	return db, scanapi
}

func NewPostgresDB(ctx context.Context, connString string) *pgxpool.Pool {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("unable to parse postgres connection uri: %v", err)
	}
	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("unable to create pg connection: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("unable to ping db: %v", err)
	}
	return db
}

func NewScanApi() *pgxscan.API {
	scanApi, err := pgxscan.NewDBScanAPI(dbscan.WithAllowUnknownColumns(true))
	if err != nil {
		log.Fatalf("failed to initial db-scan-api: %+v", err)
	}

	api, err := pgxscan.NewAPI(scanApi)
	if err != nil {
		log.Fatalf("failed to initial pgxscan-api: %v", err)
	}
	return api
}
