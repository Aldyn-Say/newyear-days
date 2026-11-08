package countdown

import "time"

func DaysUntilNewYear(date time.Time) int {
	year, month, day := date.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	nextNewYear := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)
	return int(nextNewYear.Sub(start).Hours() / 24)
}
