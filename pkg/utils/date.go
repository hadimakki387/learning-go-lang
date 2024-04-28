package utils

import "time"

func CurrentDate() string {
	// Return current date in format "2006-01-02".
	return time.Now().Format(time.RFC3339)
}

func AddDaysToDate(days int) string {
	// Add days to current date.
	return time.Now().AddDate(0, 0, days).Format(time.RFC3339)
}

func AddMonthsToDate(months int) string {
	// Add months to current date.
	return time.Now().AddDate(0, months, 0).Format(time.RFC3339)
}

func AddYearsToDate(years int) string {
	// Add years to current date.
	return time.Now().AddDate(years, 0, 0).Format(time.RFC3339)
}

func AddHoursToDate(hours int) string {

	return time.Now().Add(time.Duration(hours) * time.Hour).Format(time.RFC3339)
}

func AddMinutesToDate(minutes int) string {

	return time.Now().Add(time.Duration(minutes) * time.Minute).Format(time.RFC3339)
}

func AddSecondsToDate(seconds int) string {

	return time.Now().Add(time.Duration(seconds) * time.Second).Format(time.RFC3339)
}

func AnyFunction() {
	print("Any function.")
	// Any function.
}
