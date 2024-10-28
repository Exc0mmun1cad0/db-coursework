package models

import (
	"time"
)

type Book struct {
	ID            uint64
	Title         string
	Authors       []Author
	Categories    Category
	YearPublished int64
	Publishers    []Publisher
	ISBN          string
	Amount        uint64
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
	ID          uint64
	LastName    string
	FirstName   string
	FatherName  string
	Gender      string
	DateofBirth time.Time `time_format:"api_datetime"`
	Phone       string
	Email       string
	Address     string
}

type BookLoan struct {
	ID           uint64
	Book         Book
	Customer     Customer
	DateLoaned   time.Time
	DateDue      time.Time
	DateReturned time.Time
	OverFineDue  time.Time
	Amount       uint64
}
