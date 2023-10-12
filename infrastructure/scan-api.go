package infrastructure

import (
	"log"

	"github.com/georgysavva/scany/dbscan"
	"github.com/georgysavva/scany/sqlscan"
)

func NewScanApi() *sqlscan.API {
	scanApi, err := sqlscan.NewDBScanAPI(dbscan.WithAllowUnknownColumns(true))
	if err != nil {
		log.Fatalf("error create dbscanapi: %v", err)
	}

	api, err := sqlscan.NewAPI(scanApi)
	if err != nil {
		log.Fatalf("error create sqlscanner api: %v", err)
	}
	return api
}
