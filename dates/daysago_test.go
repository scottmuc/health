package dates_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dayDifference(then time.Time, now time.Time) int {
	return now.Day() - then.Day()
}

func parseDate(date string) time.Time {
	timeFormat := "2006-01-02 15:04:05 MST"
	returnVal, err := time.Parse(timeFormat, date)
	Ω(err).Should(BeNil())
	return returnVal
}

var _ = Describe("Calculating number of days between two dates", func() {
	It("returns 0 if the dates are the same", func() {
		then := parseDate("1990-09-03 17:35:43 UTC")
		now := parseDate("1990-09-03 18:32:15 UTC")
		daysAgo := dayDifference(then, now)
		Expect(daysAgo).To(Equal(0))
	})

	It("returns 1 if the dates are a day apart", func() {
		then := parseDate("1990-09-03 17:35:43 UTC")
		now := parseDate("1990-09-04 18:32:15 UTC")
		daysAgo := dayDifference(then, now)
		Expect(daysAgo).To(Equal(1))
	})
})
