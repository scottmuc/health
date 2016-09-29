package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Running the health binary", func() {
	It("exits with a status code of 0", func() {
		testStravaAccessToken := getEnvOrFail("STRAVA_ACCESS_TOKEN")
		session := executeBin("-stravaAccessToken", testStravaAccessToken)
		Eventually(session).Should(Exit(0))
	})

	It("prints out an acvitity and when it happened", func() {
		testStravaAccessToken := getEnvOrFail("STRAVA_ACCESS_TOKEN")
		session := executeBin("-stravaAccessToken", testStravaAccessToken)
		Eventually(session).Should(Say(`.+\. Occurred [0-9]+ day\(s\) ago`))
	})

	Context("when strava access token is not provided", func() {
		It("it writes a helpful message about the missing access token", func() {
			session := executeBin()
			Eventually(session).Should(Exit(0))
			Eventually(session.Out).Should(Say("No stravaAccessToken provided, ignoring..."))
		})
	})
})
