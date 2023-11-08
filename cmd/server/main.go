package main

import (
	"context"
	"fmt"
	"time"

	"github.com/centraldigital/cfw-core-lib/pkg/configuration/logger"

	"labForBosz/infrastructure"
	cusservice "labForBosz/internal/core/cus-service"
	"labForBosz/internal/core/domain"
	cusrepo "labForBosz/internal/repository/cus-repo"
	"labForBosz/property"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	property.Init(ctx)

	log := logger.GetLogger()
	defer log.Sync()

	// init infrastructure
	db := infrastructure.NewSqlDB()

	cusRepo := cusrepo.New(db)
	cusSvc := cusservice.New(cusRepo)

	ti := time.Now()
	id, err := cusSvc.CreateCustomerAddressTransaction(ctx, domain.CreateCustomerAddress{
		Customer: domain.Customer{
			FirstName:    "test_tx_name",
			LastName:     "test_tx_last-name",
			MobileNo:     "test_tx_mobile-no",
			The1MemberId: nil,
			The1MobileNo: nil,
			IsActive:     true,
			CreatedAt:    ti,
			CreatedBy:    "me",
			UpdatedAt:    ti,
			UpdatedBy:    "me",
		},
		Address: domain.Address{
			Subdistrict: "test_sub_district",
			District:    "test_district",
			Province:    "test_province",
		},
	})
	if err != nil {
		fmt.Printf("\nid: %v\n\n", id)
		fmt.Printf("\nerr: %v\n\n", err)
	} else {
		fmt.Printf("\nid: %v\n\n", *id)
		fmt.Printf("\nerr: %v\n\n", err)
	}

	// db, err := dbctx.OpenPgxPool(ctx, property.Get().DB.PostgresConnectionUri)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// scanapi := infrastructure.NewScanApi()

	// // init repository
	// customerRepo := customerrepo.New(db, scanapi)
	// _ = customerRepo

	// // init service
	// cusService := service.New(customerRepo)

	// // init handler
	// hdl := handler.New(cusService)

	// // create new gin engine
	// // engine := setupServer.InitServer()
	// engine := gin.New()

	// // middleware
	// // engine.Use(dbctx.GinMiddleware(db))

	// // setup router
	// router.Setup(engine, hdl)

	// setupServer.StartServer(engine, log, "localhost", "8080")

}
