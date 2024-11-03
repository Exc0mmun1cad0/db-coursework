package dateutils

import (
	"math/rand"
	"time"
)

var (
	months = map[string]int{
		"January": 1, "February": 2, "March": 3, "April": 4,
		"May": 5, "June": 6, "July": 7, "August": 8,
		"September": 9, "October": 10, "November": 11, "December": 12,
	}
)

func RandomDateAfter(date time.Time, periodInYears int) time.Time {
	d, m, y := date.Day(), date.Month(), date.Year()
	if !ValiDate(d, months[m.String()], y) {
		panic("Invalid date")
	}

	// The date after which there will be a random date
	dateSince := time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	var result int64
	if periodInYears == 0 {
		result = rand.Int63() + dateSince.Unix()
	} else {
		datePeriod := 365 * 24 * 3600 * int64(periodInYears)
		result = rand.Int63n(datePeriod) + dateSince.Unix()
	}

	return time.Unix(result, 0)
}

func FormatDate(date time.Time) string {
	tTime := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	return tTime.Format(`'2006-01-02'`)
}
