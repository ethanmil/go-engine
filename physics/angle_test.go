package physics

import (
	"math"
	"testing"
)

func TestNewAngle(t *testing.T) {
	var a = NewAngle(1)

	if !(a == Angle(1)) {
		t.Error("Error creating an angle", a)
	}
}

func TestGetVector(t *testing.T) {
	var a1 = NewAngle(math.Pi / 4)
	if !(floatsEqual(a1.GetVector().x, math.Sqrt(2)/2)) {
		t.Error("Error deriving vector's x value", a1.GetVector().x)
	}

	if !(floatsEqual(a1.GetVector().y, math.Sqrt(2)/2)) {
		t.Error("Error deriving vector's y value", a1.GetVector().y)
	}

	var a2 = NewAngle(math.Pi / 2)
	if !(floatsEqual(a2.GetVector().x, 0)) {
		t.Error("Error deriving vector's x value", a2.GetVector().x)
	}

	if !(floatsEqual(a2.GetVector().y, 1)) {
		t.Error("Error deriving vector's y value", a2.GetVector().y)
	}
}

func TestGetAngleDegrees(t *testing.T) {
	var a = NewAngle(math.Pi / 2)

	if !(floatsEqual(a.GetDegrees(), 90)) {
		t.Error("Error deriving angle's degree value", a.GetDegrees())
	}
}

func TestGetDegrees(t *testing.T) {
	if !(floatsEqual(getDegrees(math.Pi/2), 90)) {
		t.Error("Error deriving vector's x value", getDegrees(math.Pi/2))
	}
}

func TestGetRadians(t *testing.T) {
	if !(floatsEqual(getRadians(90), math.Pi/2)) {
		t.Error("Error deriving vector's x value", getRadians(90))
	}
}
