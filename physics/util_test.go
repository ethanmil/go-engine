package physics_test

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ethanmil/go-engine/physics"
)

var _ = Describe("Util", func() {
	Context("GetDegrees", func() {
		It("succeeds", func() {
			Expect(GetDegrees(math.Pi / 2)).Should(BeNumerically("~", 90))
		})
	})

	Context("GetRadians", func() {
		It("succeeds", func() {
			Expect(GetRadians(90)).Should(BeNumerically("~", math.Pi/2))
		})
	})
})
