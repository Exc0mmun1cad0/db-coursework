package main

import (
	"db-coursework/internal/api/chitaigorod"
	"db-coursework/internal/api/randomdatatools"
	"db-coursework/internal/config"
	"db-coursework/internal/mapping"
	"db-coursework/internal/repositories/books"
	"db-coursework/internal/repositories/customers"
	"db-coursework/lib/randombookloan"
	"db-coursework/pkg/postgresql"
	"fmt"
	"log"
	"time"
)

var (
	customersCount = 100
	booksCount     = 41
	bookloansCount = 50
	dateSince      = time.Date(2020, time.Month(5), 7, 0, 0, 0, 0, time.Local)
)

func main() {
	generateData()
}

func generateData() {
	// postgres initalization
	cfg := config.MustLoad()
	conn, err := postgresql.NewClient(cfg.PostgreSQL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("initialized db for storing library data")

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
	log.Println("customers data was added")

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

	bookLoans := randombookloan.GenerateBookLoans(bookloansCount, dateSince, uint64(len(customerIDs)), uint64(len(bookIDs)))
	log.Println("generated bookloans")
	fmt.Println(bookLoans)
	for _, bookLoan := range bookLoans {
		fmt.Println(bookLoan)
		_, err := bookRepo.AddBookLoan(bookLoan)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("bookloans data was added")
}
