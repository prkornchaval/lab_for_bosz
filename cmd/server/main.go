package main

import (
	"context"

	setupServer "github.com/centraldigital/cfw-core-lib/pkg/backoffice/setup-server"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/logger"
	"github.com/gin-gonic/gin"

	"labForBosz/infrastructure"
	"labForBosz/internal/core/service"
	"labForBosz/internal/handler"
	customerrepo "labForBosz/internal/repository/customer-repository"
	"labForBosz/internal/router"
	"labForBosz/pkg/dbctx"
	"labForBosz/property"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	property.Init(ctx)

	log := logger.GetLogger()
	defer log.Sync()

	// init infrastructure
	db, err := dbctx.OpenPgxPool(ctx, property.Get().DB.PostgresConnectionUri)
	if err != nil {
		log.Fatal(err)
	}
	scanapi := infrastructure.NewScanApi()

	// init repository
	customerRepo := customerrepo.New(db, scanapi)
	_ = customerRepo

	// init service
	cusService := service.New(customerRepo)

	// init handler
	hdl := handler.New(cusService)

	// create new gin engine
	// engine := setupServer.InitServer()
	engine := gin.New()

	// middleware
	// engine.Use(dbctx.GinMiddleware(db))

	// setup router
	router.Setup(engine, hdl)

	setupServer.StartServer(engine, log, "localhost", "8080")

}
