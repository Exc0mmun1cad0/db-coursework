package models

import "time"

type Book struct {
	ID         uint64
	Title      string
	Authors    []Author
	Categories []Category
	Year       int64
	Publishers []Publisher
	ISBN       string
}

type Author struct {
	ID   uint64
	Name string
}

type Category struct {
	ID   uint64
	Name string
}

type Publisher struct {
	ID   uint64
	Name string
}

type Customer struct {
	LastName   string
	FirstName  string
	FatherName string
	Gender     string
	BirthDate  time.Time
	Phone      string
	Email      string
	Address    string
}

type BookLoan struct {
	ID           uint64
	BookID       uint64
	CustomerID   uint64
	DateLoaned   time.Time
	DateDue      time.Time
	DateReturned time.Time
	OverFineDue  time.Time
	Amount       uint64
}
