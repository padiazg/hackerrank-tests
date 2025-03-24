package main

func isLeapYearJulian(year int32) bool {
	return (year % 4) == 0
}

func isLeapYearGregorian(year int32) bool {
	return ((year % 400) == 0) || ((year%4) == 0) && ((year%100) != 0)
}

func generator(year int32) func() (int, int32) {
	var (
		months = map[int]int32{1: 31, 2: 28, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31}
		month  = 0
	)

	if (year < 1918 && isLeapYearJulian(year)) || (year > 1918 && isLeapYearGregorian(year)) {
		months[2]++
	}

	if year == 1918 {
		months[2] -= 13
	}

	return func() (int, int32) {
		if month++; month > 12 {
			return 0, 0
		}
		return month, months[month]
	}
}

// next := generator(year)
// target := int32(256)
// for {
// 	month, daysInMonth := next()
// 	if month == 0 {
// 		break
// 	}

// 	if target>daysInMonth {
// 		target -= daysInMonth
// 	}

// }
