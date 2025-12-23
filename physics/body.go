package physics

import (
	"fmt"
)

// Axis aligned rectangles
type Body struct {
	Position, Size Vector2f
}

func NewBody(position, size Vector2f) Body {
	return Body{position, size}
}

// Returns the center of the body
func (body Body) Center() Vector2f {
	var center Vector2f

	// The center is the sum of the position and half the size of the body
	center.X = body.Position.X + body.Size.X/2
	center.Y = body.Position.Y + body.Size.Y/2

	return center
}

// Pretty formatting
func (body Body) String() string {
	return fmt.Sprintf("%f\t%f\t\n%f\t%f\t", body.Position.X, body.Position.Y, body.Size.X, body.Size.Y)
}
