package infrastructure

import (
	"context"
	"labForBosz/property"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(ctx context.Context) *pgxpool.Pool {
	connString := property.Get().DB.PostgresConnectionUri
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

/*
import (
	"context"
	"fmt"
	"net/url"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/logger"
	"github.com/centraldigital/cfw-core-order-jobs/order-pending-kbank-payment/property"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context) *pgxpool.Pool {
	log := logger.ExtractLogger(ctx)

	password := url.QueryEscape(property.Get().Postgres.Password)
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		property.Get().Postgres.User,
		password,
		property.Get().Postgres.Host,
		property.Get().Postgres.Port,
		property.Get().Postgres.Database,
	)
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("unable to parse postgres connection uri: %v", err)
	}

	cfg.ConnConfig.Tracer = otelpgx.NewTracer()
	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("unable to create pg connection: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("unable to ping db: %v", err)
	}

	return db
}

*/
