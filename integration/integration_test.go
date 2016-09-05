package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Running the health binary", func() {
	var pathToHealthCLI string

	BeforeSuite(func() {
		var err error
		pathToHealthCLI, err = gexec.Build("github.com/scottmuc/health")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("exits with a status code of 0", func() {
		testStravaAccessToken := getEnvOrFail("STRAVA_ACCESS_TOKEN")
		command := exec.Command(pathToHealthCLI, "-stravaAccessToken", testStravaAccessToken)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
	})

	It("prints out an acvitity and when it happened", func() {
		testStravaAccessToken := getEnvOrFail("STRAVA_ACCESS_TOKEN")
		command := exec.Command(pathToHealthCLI, "-stravaAccessToken", testStravaAccessToken)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gbytes.Say(`.+\. Occured on .+`))
	})

	Context("when strava access token is not provided", func() {
		It("does not exit 0", func() {
			command := exec.Command(pathToHealthCLI)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).ShouldNot(gexec.Exit(0))
		})

		It("it writes a helpful message about the missing access token", func() {
			command := exec.Command(pathToHealthCLI)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session.Err).Should(gbytes.Say("stravaAccessToken is required"))
		})
	})
})
