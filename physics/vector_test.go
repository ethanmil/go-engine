package physics

import (
	"math"
	"testing"
)

func TestNewVector(t *testing.T) {
	var v = NewVector(1, 2)

	if !(v == Vector{
		x: 1,
		y: 2,
	}) {
		t.Error("Error creating a vector", v)
	}
}

func TestGetAngle(t *testing.T) {
	var v1 = NewVector(1, 1)

	if !(v1.GetAngle() == float64(math.Pi/4)) {
		t.Error("Error calculating angle", v1.GetAngle())
	}

	var v2 = NewVector(0, 0)

	if !(v2.GetAngle() == float64(0)) {
		t.Error("Error calculating angle", v2.GetAngle())
	}
}
