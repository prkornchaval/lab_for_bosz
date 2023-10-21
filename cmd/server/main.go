package main

import (
	"context"

	setupServer "github.com/centraldigital/cfw-core-lib/pkg/backoffice/setup-server"
	"github.com/centraldigital/cfw-core-lib/pkg/configuration/logger"
	"github.com/gin-gonic/gin"

	"labForBosz/internal/core/service"
	"labForBosz/internal/handler"
	customerrepo "labForBosz/internal/repository/customer-repository"
	"labForBosz/internal/router"
	"labForBosz/pkg/dbctx"
	"labForBosz/pkg/infrastructure"
	"labForBosz/property"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	property.Init(ctx)

	log := logger.GetLogger()
	defer log.Sync()

	// init infrastructure
	db, scanapi := infrastructure.NewPostgresWithScanApi(ctx, property.Get().DB.PostgresConnectionUri)
	_, _ = db, scanapi

	// init repository
	customerRepo := customerrepo.New(scanapi)
	_ = customerRepo

	// init service
	cusService := service.New(customerRepo)

	// init handler
	hdl := handler.New(cusService)

	// create new gin engine
	// engine := setupServer.InitServer()
	engine := gin.New()

	// middleware
	engine.Use(dbctx.GinMiddleware(db))

	// setup router
	router.Setup(engine, hdl)

	setupServer.StartServer(engine, log, "localhost", "8080")

}
