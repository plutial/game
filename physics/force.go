package physics

import (
	"math"
)

// Stores an x and a y
type Vector2 struct {
	X, Y float64
}

func NewVector2(x, y float64) Vector2 {
	return Vector2{x, y}
}

// Collisions are going to be inside force
type Collisions struct {
	Left, Right, Up, Down bool
}

// Movement and collisions
type Force struct {
	// Instantaneous movement
	Velocity Vector2

	// Persisting momentum
	Acceleration Vector2

	// Maximum horizontal speed
	Speed float64

	// Collisions
	Collisions Collisions
}

func NewForce(velocity, acceleration Vector2) Force {
	force := Force{}

	force.Velocity = velocity
	force.Acceleration = acceleration

	force.Speed = 3

	force.Collisions = Collisions{}

	return force
}

// Just pythagoras
func (a *Vector2) GetDistance(b Vector2) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

// Rounding
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Update the collisions
func (collisions *Collisions) Update(contactNormal Vector2) {
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
