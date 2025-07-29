package physics

import (
	// raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
	Position, Size rl.Vector2
}

func NewBody(position rl.Vector2, size rl.Vector2) Body {
	return Body{position, size}
}

func (body *Body) Rectangle() rl.Rectangle {
	rectangle := rl.NewRectangle(body.Position.X, body.Position.Y, body.Size.X, body.Size.Y)

	return rectangle
}

func (body *Body) Center() rl.Vector2 {
	var center rl.Vector2

	// The center is the sum of the position and half the size of the body
	center.X = body.Position.X + body.Size.X/2
	center.Y = body.Position.Y + body.Size.Y/2

	return center
}
