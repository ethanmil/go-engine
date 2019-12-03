package animation

import (
	"time"

	"github.com/ethanmil/go-engine/physics"
	"github.com/veandco/go-sdl2/sdl"
)

// Sprite -
type Sprite struct {
	Size  physics.Vector
	Chunk sdl.Rect
	Angle float64

	// log helper
	LastLogged time.Time
}

// Draw -
func (s *Sprite) Draw(position physics.Vector, angle float64, texture *sdl.Texture, renderer *sdl.Renderer) {
	renderer.CopyEx(
		texture,
		&sdl.Rect{X: int32(s.Chunk.X), Y: int32(s.Chunk.Y), W: int32(s.Size.X), H: int32(s.Size.Y)},
		&sdl.Rect{X: int32(position.X), Y: int32(position.Y), W: int32(s.Chunk.W), H: int32(s.Chunk.H)},
		angle,
		&sdl.Point{X: int32(s.Size.X / 2), Y: int32(s.Size.Y / 2)},
		sdl.FLIP_NONE,
	)
}
