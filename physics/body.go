package physics

import (
	// raylib
    rl "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
    Position, Size rl.Vector2
}

func NewBody(position rl.Vector2, size rl.Vector2) Body {
	return Body {position, size}
}
