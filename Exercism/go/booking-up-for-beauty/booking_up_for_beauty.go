package booking

import (
	"time"
)

// data especial: Mon Jan 2 15:04:05 -0700 MST 2006
// layout example: Mon, 01/02/2006, 15:04

// 7/25/2019 13:45:00
// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedDate, _ := time.Parse("1/2/2006 15:04:05", date)

	return parsedDate
}

// July 25, 2019 13:45:00
// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsedDate, _ := time.Parse("January 2, 2006 15:04:05", date)

	return parsedDate.Before(time.Now())
}

// Thursday, July 25, 2019 13:45:00
// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return parsedDate.Hour() >= 12 && parsedDate.Hour() < 18
}

// You have an appointment on Monday, June 06, 2005, at 10:30.
// You have an appointment on Monday, June 6, 2005, at 10:30.

// 7/25/2019 13:45:00
// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedDate, _ := time.Parse("1/2/2006 15:04:05", date)
	formattedDate := parsedDate.Format("Monday, January 2, 2006, at 15:04.")

	return "You have an appointment on " + formattedDate
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
