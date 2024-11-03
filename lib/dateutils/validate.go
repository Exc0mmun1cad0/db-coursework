package dateutils

import "slices"

func ValiDate(day, month, year int) bool {
	if year < 1 || day < 1 || month < 1 {
		return false
	}

	switch {
	case slices.Contains([]int{1, 3, 5, 7, 8, 10, 12}, month):
		if day <= 31 {
			return true
		}
	case slices.Contains([]int{4, 6, 9, 11}, month):
		if day <= 30 {
			return true
		}
	case month == 2:
		i := 0
		if IsLeapYear(year) {
			i = 1
		}
		if day <= 28+i {
			return true
		}
	}

	return false
}

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return false
			} else {
				return true
			}
		} else {
			return true
		}
	}

	return false
}
