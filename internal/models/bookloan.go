package models

import "time"

type BookLoan struct {
	ID           uint64
	Book         uint64
	Customer     uint64
	DateLoaned   time.Time
	DateDue      time.Time
	DateReturned time.Time
	Amount       uint64
}
