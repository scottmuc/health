package dates_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDates(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dates Suite")
}
