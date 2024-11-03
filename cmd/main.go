package main

import (
	"db-coursework/internal/api/chitaigorod"
	"db-coursework/internal/api/randomdatatools"
	"db-coursework/internal/config"
	"db-coursework/internal/mapping"
	"db-coursework/internal/repositories/books"
	"db-coursework/internal/repositories/customers"
	"db-coursework/pkg/postgresql"
	"fmt"
	"log"
)

// TODO: move to config
var (
	customersCount = 100
	booksCount     = 41
)

func main() {
	// postgres initalization
	cfg := config.MustLoad()
	conn, err := postgresql.NewClient(cfg.PostgreSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("initialized db for storing library data")

	// doing migrations
	err = postgresql.Migrate(cfg.PostgreSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Adding customers
	customerRepo := customers.NewRepository(conn)
	customers, err := randomdatatools.GetCustomers(customersCount)
	if err != nil {
		log.Fatal(err)
	}
	customerIDs, err := customerRepo.AddCustomers(customers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("customers data was added")

	// Adding books and related objects
	bookRepo := books.NewRepository(conn)
	c, err := chitaigorod.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	responseBooks, err := c.GetBooks(booksCount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("received books data")
	books, err := mapping.ResponseToModel(responseBooks)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("mapping completed")
	log.Println(books[0].Category.ID, books[1].Category.ID)
	bookIDs, err := bookRepo.AddBooks(books)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("books data was added")

}
