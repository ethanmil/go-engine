package physics

var epsilon float64 = 0.000001

func floatsEqual(a, b float64) bool {
	if (a-b) < epsilon && (b-a) < epsilon {
		return true
	}
	return false
}
