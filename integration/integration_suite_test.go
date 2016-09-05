package integration_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func getEnvOrFail(name string) string {
	envVar := os.Getenv(name)
	Expect(envVar).ToNot(BeEmpty(), fmt.Sprintf("Expected the environment variable '%s' to be set", name))
	return envVar
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
