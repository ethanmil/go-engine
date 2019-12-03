package physics

import "math"

type angle float64

func newAngle(radians float64) angle {
	return angle(radians)
}

func (a angle) reset() {
	a = angle(0)
}

func (a angle) getVector() (v vector) {
	v.x = math.Cos(float64(a))
	v.y = math.Sin(float64(a))
	return v
}

func (a angle) getDegrees() float64 {
	return getDegrees(float64(a))
}

func getDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

func getRadians(degrees float64) float64 {
	return degrees / (180 / math.Pi)
}
