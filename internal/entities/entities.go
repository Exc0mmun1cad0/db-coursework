package entities

import "time"

type Book struct {
	ID            uint64
	Title         string
	CategoryID    uint64
	YearPublished int64
	ISBN          string
	Amount        uint64
}

type Author struct {
	ID   uint64
	Name string
}

type BookAuthor struct {
	BookID   uint64
	AuthorID uint64
}

type Publisher struct {
	ID   uint64
	Name string
}

type BookPublisher struct {
	BookID   uint64
	AuthorID uint64
}

type Category struct {
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
