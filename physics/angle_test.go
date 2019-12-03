package physics

import (
	"math"
	"testing"
)

func TestNewAngle(t *testing.T) {
	var a = newAngle(1)

	if !(a == angle(1)) {
		t.Error("Error creating an angle", a)
	}
}

func TestGetVector(t *testing.T) {
	var a1 = newAngle(math.Pi / 4)
	if !(floatsEqual(a1.getVector().x, math.Sqrt(2)/2)) {
		t.Error("Error deriving vector's x value", a1.getVector().x)
	}

	if !(floatsEqual(a1.getVector().y, math.Sqrt(2)/2)) {
		t.Error("Error deriving vector's y value", a1.getVector().y)
	}

	var a2 = newAngle(math.Pi / 2)
	if !(floatsEqual(a2.getVector().x, 0)) {
		t.Error("Error deriving vector's x value", a2.getVector().x)
	}

	if !(floatsEqual(a2.getVector().y, 1)) {
		t.Error("Error deriving vector's y value", a2.getVector().y)
	}
}

func TestGetAngleDegrees(t *testing.T) {
	var a = newAngle(math.Pi / 2)

	if !(floatsEqual(a.getDegrees(), 90)) {
		t.Error("Error deriving angle's degree value", a.getDegrees())
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
