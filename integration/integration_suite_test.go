package integration_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

func getEnvOrFail(name string) string {
	envVar := os.Getenv(name)
	Expect(envVar).ToNot(BeEmpty())
	return envVar
}
