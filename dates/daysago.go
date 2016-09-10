package dates

import "time"

// DayDifference returns the number of days ago it is between two dates. The actual amount of time
// doesn't really matter, it's the calendar date that this function uses. So if you have an activity
// at 11:59pm yesterday, and current time is the next day at 1 minute after minute, this function will
// return 1 because it happened a day ago.
func DayDifference(then time.Time, now time.Time) int {
	normalizedThen := time.Date(then.Year(), then.Month(), then.Day(), 23, 59, 59, 0, time.UTC)
	normalizedNow := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
	hoursAgo := normalizedNow.Sub(normalizedThen).Hours()
	// not sure if I am potentially introducing rounding issues by not casting hoursAgo to an
	// int before doing the division.
	return int(hoursAgo / 24)
}
