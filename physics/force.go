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
	// Instantaneous movement
	Velocity rl.Vector2

	// Persisting momentum
	Acceleration rl.Vector2

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

func (force *Force) UpdateGravity() {
	// Apply gravity
	force.Acceleration.Y += 0.3

	// Limit the gravity to 5
	force.Acceleration.Y = min(5, force.Acceleration.Y)

	// If the body is on the ground, lower the gravity
	// Don't set it to zero, because, then, the entity is flying
	if force.Collisions.Down {
		force.Acceleration.Y = min(0.3, force.Acceleration.Y)
	}
}
