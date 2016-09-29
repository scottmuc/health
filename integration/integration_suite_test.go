package integration_test

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var pathToHealthCLI string

var _ = BeforeSuite(func() {
	var err error
	pathToHealthCLI, err = gexec.Build("github.com/scottmuc/health")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func getEnvOrFail(name string) string {
	envVar := os.Getenv(name)
	Expect(envVar).ToNot(BeEmpty(), fmt.Sprintf("Expected the environment variable '%s' to be set", name))
	return envVar
}

func executeBin(args ...string) *gexec.Session {
	command := exec.Command(pathToHealthCLI, args...)
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}
