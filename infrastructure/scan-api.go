package infrastructure

import (
	"log"

	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/georgysavva/scany/v2/pgxscan"
)

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
