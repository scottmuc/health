package dates_test

import (
	"time"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/scottmuc/health/dates"
)

func parseDate(date string) time.Time {
	timeFormat := "2006-01-02 15:04:05 MST"
	returnVal, err := time.Parse(timeFormat, date)
	Ω(err).Should(BeNil())
	return returnVal
}

var _ = DescribeTable("DayDifference",
	func(then string, now string, expected int) {
		thenDate := parseDate(then)
		nowDate := parseDate(now)
		daysAgo := dates.DayDifference(thenDate, nowDate)
		Expect(daysAgo).To(Equal(expected))
	},
	Entry("same day", "1990-09-03 17:35:43 UTC", "1990-09-03 18:32:15 UTC", 0),
	Entry("same day boundaries", "1990-09-03 00:00:00 UTC", "1990-09-03 23:59:59 UTC", 0),
	Entry("day boundary", "1990-09-03 23:35:43 UTC", "1990-09-04 00:32:15 UTC", 1),
	Entry("next day extremes", "1990-09-03 00:00:00 UTC", "1990-09-04 23:59:59 UTC", 1),
	Entry("month boundary", "1990-09-30 17:35:43 UTC", "1990-10-01 18:32:15 UTC", 1),
	Entry("year boundary", "1990-12-31 23:35:43 UTC", "1991-01-01 00:32:15 UTC", 1),
	Entry("leap year", "2000-02-28 23:35:43 UTC", "2000-03-01 00:32:15 UTC", 2),
)
