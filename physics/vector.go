package physics

import "math"

type vector struct {
	x float64
	y float64
}

func newVector(x, y float64) vector {
	return vector{
		x: x,
		y: y,
	}
}

func (v *vector) reset() {
	v.x = 0
	v.y = 0
}

func (v *vector) getAngle() float64 {
	return math.Atan2(v.y, v.x)
}
