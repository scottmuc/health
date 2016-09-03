package integration_test

import (
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var (
	pathToHealthCLI       string
	testStravaAccessToken string
)

func getEnvOrLeaveUnset(name string) string {
	tmp := os.Getenv(name)
	if tmp == "" {
		tmp = "not set"
	}
	return tmp
}

var _ = Describe("Integration", func() {
	BeforeSuite(func() {
		var err error
		testStravaAccessToken = getEnvOrLeaveUnset("STRAVA_ACCESS_TOKEN")
		pathToHealthCLI, err = gexec.Build("github.com/scottmuc/health")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("exits with a status code of 0", func() {
		Expect(testStravaAccessToken).ToNot(Equal("not set"))
		command := exec.Command(pathToHealthCLI, "-stravaAccessToken", testStravaAccessToken)
		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
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
			Eventually(session).Should(gbytes.Say("stravaAccessToken is required"))
		})
	})
})
