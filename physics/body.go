package physics

type Body struct {
	Position, Size Vector2
}

func NewBody(position, size Vector2) Body {
	return Body{position, size}
}

func (body *Body) Center() Vector2 {
	var center Vector2

	// The center is the sum of the position and half the size of the body
	center.X = body.Position.X + body.Size.X/2
	center.Y = body.Position.Y + body.Size.Y/2

	return center
}
