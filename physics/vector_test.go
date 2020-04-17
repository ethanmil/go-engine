package physics_test

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ethanmil/go-engine/physics"
)

var _ = Describe("Vector", func() {
	Context("Builds", func() {
		It("succeeds", func() {
			vector := NewVector(1, 2)

			Expect(vector).To(Equal(Vector{
				X: 1,
				Y: 2,
			}))
		})
	})

	Context("GetAngle", func() {
		It("1, 1", func() {
			vector := NewVector(1, 1)

			Expect(vector.GetAngle()).Should(BeNumerically("~", math.Pi/4))
		})

		It("0, 0", func() {
			vector := NewVector(0, 0)

			Expect(vector.GetAngle()).Should(BeNumerically("~", 0))
		})
	})

	Context("Reset", func() {
		It("should succeed", func() {
			vector := NewVector(1, 1)

			vector.Reset()

			Expect(vector).To(Equal(Vector{
				X: 0,
				Y: 0,
			}))
		})
	})
})
