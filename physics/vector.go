package physics

import "math"

// Vector -
type Vector struct {
	x float64
	y float64
}

// NewVector -
func NewVector(x, y float64) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

// Reset -
func (v *Vector) Reset() {
	v.x = 0
	v.y = 0
}

// GetAngle -
func (v *Vector) GetAngle() float64 {
	return math.Atan2(v.y, v.x)
}
