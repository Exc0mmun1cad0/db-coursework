package main

import (
	rdt "db-coursework/internal/api/random_data_tools"
	"db-coursework/internal/config"
	"db-coursework/pkg/postgresql"
	"fmt"
	"log"

	"db-coursework/internal/repositories/customers"
)

// TODO: create usecase with repository interfaces as fields and pass their implementations to it
func main() {
	cfg := config.MustLoad()

	conn, err := postgresql.NewClient(cfg.PostgreSQL)
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err)
	}
	_ = conn

	postgresql.Migrate(cfg.PostgreSQL)

	customersData, err := rdt.GetCustomers(10)
	if err != nil {
		log.Fatal(err)
	}

	customersRepo := customers.NewRepository(conn)

	_, err = customersRepo.AddCustomers(customersData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("added customers data")

}
