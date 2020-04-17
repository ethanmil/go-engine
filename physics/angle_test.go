package physics

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Angle", func() {

	Context("NewAngle", func() {
		It("should build", func() {
			angle := NewAngle(1)

			Expect(angle).To(Equal(Angle(1)))
		})
	})

	Context("GetVector", func() {
		It("45 degrees", func() {
			angle := NewAngle(math.Pi / 4)
			vector := angle.GetVector()

			Expect(vector.x).Should(BeNumerically("~", math.Sqrt(2)/2))
			Expect(vector.y).Should(BeNumerically("~", math.Sqrt(2)/2))
		})

		It("90 degrees", func() {
			angle := NewAngle(math.Pi / 2)
			vector := angle.GetVector()

			Expect(vector.x).Should(BeNumerically("~", 0))
			Expect(vector.y).Should(BeNumerically("~", 1))
		})
	})

	Context("GetAngleDegrees", func() {
		It("90 degrees", func() {
			angle := NewAngle(math.Pi / 2)

			Expect(angle.GetDegrees()).Should(BeNumerically("~", 90))
		})
	})
})
