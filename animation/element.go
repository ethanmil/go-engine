package animation

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/ethanmil/go-engine/physics"
)

// Element -
type Element struct {
	Sprite         Sprite
	Position       physics.Vector
	Speed          float64
	Angle          physics.Angle
	CollisionSpace []int
	Active         bool

	// log helper
	lastLogged time.Time
}

// Draw -
func (e *Element) Draw(texture *sdl.Texture, renderer *sdl.Renderer) {
	if !e.Active {
		return
	}

	e.Sprite.Draw(e.Position, e.Angle.GetDegrees(), texture, renderer)
}

// Update -
func (e *Element) Update(delta float64) {
	movement := e.Angle.GetVector()
	e.Position.X += movement.X * e.Speed * delta
	e.Position.Y += movement.Y * e.Speed * delta
}

// Print -
func (e *Element) Print(every time.Duration) {
	if time.Since(e.lastLogged) >= every {
		e.lastLogged = time.Now()
		println(fmt.Sprintf("element: %+v", e))
	}
}
