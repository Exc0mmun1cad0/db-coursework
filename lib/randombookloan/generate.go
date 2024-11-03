// Package for models.Bookloan generation
package randombookloan

import (
	"db-coursework/internal/models"
	"db-coursework/lib/dateutils"
	"time"

	"math/rand"
)

var (
	periodInYears        = 2
	maxLoanPeriodInDays  = 60
	maxDaysUntilReturned = 60
	maxAmount            = 6
)

// Function for generating collection of book loans in the amount of count which was made since dateSince
func GenerateBookLoans(count int, dateSince time.Time, customerIDs uint64, bookIDs uint64) []models.BookLoan {
	customerSet := make(map[uint64]struct{}, 0)
	bookSet := make(map[uint64]struct{}, 0)

	result := make([]models.BookLoan, 0, count)
	for i := 0; i < count; i++ {
		dateLoaned := dateutils.RandomDateAfter(dateSince, periodInYears)
		dateDue := dateLoaned.AddDate(0, 0, rand.Intn(maxLoanPeriodInDays))
		dateReturned := dateLoaned.AddDate(0, 0, rand.Intn(maxDaysUntilReturned))
		amount := rand.Intn(maxAmount) + 1

		var customerID uint64
		customerID = rand.Uint64() % uint64(customerIDs)
		// if customerID alreayd in use, let's randomly choose one more time
		if _, ok := customerSet[customerID]; ok {
			customerID = rand.Uint64() % uint64(customerIDs)
		}

		var bookID uint64
		bookID = rand.Uint64() % uint64(bookIDs)
		// if bookID alreayd in use, let's randomly choose one more time
		if _, ok := bookSet[bookID]; ok {
			bookID = rand.Uint64() % uint64(bookIDs)
		}

		bookSet[bookID] = struct{}{}
		customerSet[customerID] = struct{}{}

		result = append(
			result,
			models.BookLoan{
				Book:         bookID + 1,
				Customer:     customerID + 1,
				DateDue:      dateDue,
				DateLoaned:   dateLoaned,
				DateReturned: dateReturned,
				Amount:       uint64(amount),
			},
		)
	}

	return result
}
