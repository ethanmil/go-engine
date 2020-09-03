package animation

import (
	"github.com/ethanmil/go-engine/physics"
	"github.com/hajimehoshi/ebiten"
)

// Element -
type Element struct {
	Sprite        *ebiten.Image
	Angle         physics.Angle
	Position      physics.Vector
	Collision     []int
	isHighlighted bool
}

// Draw -
func (e *Element) Draw(screen *ebiten.Image) {
	w, h := e.Sprite.Size()

	op := &ebiten.DrawImageOptions{}
	if e.isHighlighted {
		op.ColorM.Scale(255, 255, 0, 0.8)
	}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(e.Angle))
	op.GeoM.Translate(e.Position.X, e.Position.Y)

	screen.DrawImage(e.Sprite, op)
}

// Update -
func (e *Element) Update(speed, delta float64, overrideVector *physics.Vector) {
	movement := e.Angle.GetVector()
	if overrideVector != nil {
		movement = *overrideVector // for handling collisions
	}
	e.Position.X += movement.X * speed * delta
	e.Position.Y += movement.Y * speed * delta
}

// Highlight -
func (e *Element) Highlight() {
	e.isHighlighted = true
}

// DoesCollide -
func (e *Element) DoesCollide(nums ...int) bool {
	for _, n1 := range nums {
		for _, n2 := range e.Collision {
			if n1 == n2 {
				return true
			}
		}
	}
	return false
}
