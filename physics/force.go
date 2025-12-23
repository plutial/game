package physics

import (
	"math"
)

// Collisions are going to be inside force
type Collisions struct {
	Left, Right, Up, Down bool
}

// Movement and collisions
type Force struct {
	// Instantaneous movement
	Velocity Vector2f

	// Persisting momentum
	Acceleration Vector2f

	// Maximum horizontal speed
	Speed float64

	// Collisions
	Collisions Collisions
}

// Creates a new force
func NewForce(velocity, acceleration Vector2f) Force {
	force := Force{}

	force.Velocity = velocity
	force.Acceleration = acceleration

	force.Speed = 3

	force.Collisions = Collisions{}

	return force
}

// Rounding
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Update the collisions
func (collisions *Collisions) Update(contactNormal Vector2f) {
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

func (collisions *Collisions) Collided() bool {
	return collisions.Left || collisions.Right || collisions.Up || collisions.Down
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
