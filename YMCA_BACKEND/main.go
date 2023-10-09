package main

import (
	"YMCA_BACKEND/core"
	"YMCA_BACKEND/repository"
	"YMCA_BACKEND/service"
	"database/sql"
)

func main() {
	storage := repository.NewRepository(&sql.DB{})
	service := service.NewService(storage)
	newCoreService := core.NewCoreService(storage, service)

	newCoreService.CallRun()
}
