//* This package looks useless. Maybe I'll delete it in the future
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
	ID          uint64    `db:"id"`
	LastName    string    `db:"last_name"`
	FirstName   string    `db:"first_name"`
	FatherName  string    `db:"father_name"`
	Gender      string    `db:"gender"`
	DateOfBirth time.Time `db:"date_of_birth"`
	Phone       string    `db:"phone"`
	Email       string    `db:"email"`
	Address     string    `db:"address"`
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
