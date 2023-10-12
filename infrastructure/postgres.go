package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() *sql.DB {
	postgresConnectionURI := "host=localhost database=user-api user=postgres password=secret port=5432"
	pgxPool, err := sql.Open("pgx", postgresConnectionURI)
	if err != nil {
		log.Panicf("Failed to initialize Postgres client: %+v", err)
	}

	return pgxPool
}
