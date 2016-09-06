package strava_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/scottmuc/health/strava"
)

var _ = Describe("StravaService", func() {
	It("can be created", func() {
		service := strava.NewService("stub access token")
		Expect(service).ToNot(BeNil())
	})
})
