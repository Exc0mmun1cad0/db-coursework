package models

import "time"

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
