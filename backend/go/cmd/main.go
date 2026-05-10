package main

import (
	"fmt"
	"log"

	"github.com/iamveso/webhooks/backend/go/internal/handler"
	"github.com/iamveso/webhooks/backend/go/internal/repository"
	"github.com/iamveso/webhooks/backend/go/internal/service"
)

func main() {
	// get env variables
	//
	// initialize db
	//
	// initialize repository
	repository := repository.NewRepository()
	// initialize service
	service := service.NewService(repository)
	// start server
	handler := handler.NewHandler(service)
	if err := handler.StartServer(); err != nil {
		log.Fatalf("something went wrong when starting server: %v", err)
	}
	fmt.Println("Hello World")
}
