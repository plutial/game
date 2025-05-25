package physics

import (
	"math"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Force struct {
	Velocity, Acceleration rl.Vector2
}

func NewForce(velocity rl.Vector2, acceleration rl.Vector2) Force {
	return Force{velocity, acceleration}
}

// Just pythagoras
func GetDistance(a, b rl.Vector2) float32 {
	return float32(math.Sqrt(math.Pow(float64(a.X - b.X), 2) + math.Pow(float64(a.Y - b.Y), 2)))
}
