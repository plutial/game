package physics

import (
	"math"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Collisions are going to be inside force
type Collisions struct {
	Left, Right, Up, Down bool
}

// Movement and collisions
type Force struct {
	// Movement
	Velocity, Acceleration rl.Vector2

	// Maximum horizontal speed
	Speed float32

	// Collisions
	Collisions Collisions
}

func NewForce(velocity rl.Vector2, acceleration rl.Vector2) Force {
	force := Force{}

	force.Velocity = velocity
	force.Acceleration = acceleration

	force.Speed = 3

	force.Collisions = Collisions{}

	return force
}

// Just pythagoras
func GetDistance(a, b rl.Vector2) float32 {
	return float32(math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2)))
}

// Rounding
func Round(x, unit float32) float32 {
	return float32(math.Round(float64(x/unit))) * unit
}

// Update the collisions
func (collisions *Collisions) Update(contactNormal rl.Vector2) {
	if contactNormal.X == 1 {
		collisions.Left = true
	} else if contactNormal.X == -1 {
		collisions.Right = true
	}

	if contactNormal.Y == 1 {
		collisions.Up = true
	} else if contactNormal.Y == -1 {
		collisions.Down = true
	}
}
