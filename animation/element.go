package animation

import (
	"fmt"
	"time"

	"github.com/ethanmil/go-engine/physics"
)

type element struct {
	sprite         sprite
	position       physics.Vector
	speed          float64
	angle          physics.Angle
	collisionSpace []int
	active         bool

	// log helper
	lastLogged time.Time
}

func (e *element) draw() {
	if !e.active {
		return
	}

	e.sprite.draw(e.position, e.angle.GetDegrees(), renderer)
}

func (e *element) update() {
	movement := e.angle.GetVector()
	e.position.x += movement.x * e.speed * delta
	e.position.y += movement.y * e.speed * delta
}

func (e *element) print(every time.Duration) {
	if time.Since(e.lastLogged) >= every {
		e.lastLogged = time.Now()
		println(fmt.Sprintf("element: %+v", e))
	}
}
