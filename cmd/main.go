package main

import (
	"labForBosz/infrastructure"
	"labForBosz/internal/core/service"
	"labForBosz/internal/handler"
	"labForBosz/internal/repository"
	"labForBosz/internal/router"
)

func main() {
	// setup infrastructure
	db := infrastructure.NewDB()
	scanApi := infrastructure.NewScanApi()
	_, _ = db, scanApi

	// new repository
	repo := repository.New()

	// new service
	svc := service.New(repo)

	// new handler
	hdl := handler.New(svc)

	// new router
	r := router.NewRouter(hdl)

	// start server
	r.Run()
}
