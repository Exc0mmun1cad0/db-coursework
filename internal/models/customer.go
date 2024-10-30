package models

import (
	"time"
)

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
