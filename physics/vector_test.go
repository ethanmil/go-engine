package physics

import (
	"math"
	"testing"
)

func TestNewVector(t *testing.T) {
	var v = newVector(1, 2)

	if !(v == vector{
		x: 1,
		y: 2,
	}) {
		t.Error("Error creating a vector", v)
	}
}

func TestGetAngle(t *testing.T) {
	var v1 = newVector(1, 1)

	if !(v1.getAngle() == float64(math.Pi/4)) {
		t.Error("Error calculating angle", v1.getAngle())
	}

	var v2 = newVector(0, 0)

	if !(v2.getAngle() == float64(0)) {
		t.Error("Error calculating angle", v2.getAngle())
	}
}
