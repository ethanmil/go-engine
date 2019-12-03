package physics

import "math"

var epsilon float64 = 0.000001

func floatsEqual(a, b float64) bool {
	if (a-b) < epsilon && (b-a) < epsilon {
		return true
	}
	return false
}

func getDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}

func getRadians(degrees float64) float64 {
	return degrees / (180 / math.Pi)
}
