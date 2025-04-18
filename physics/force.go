package physics

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Force struct {
	Velocity, Acceleration rl.Vector2
}

func NewForce(velocity rl.Vector2, acceleration rl.Vector2) Force {
	return Force { velocity, acceleration }
}
