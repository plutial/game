package physics

import (
	"fmt"
	"log"
	"math"
)

const (
	// Collision Types
	CollisionNone int = iota
	CollisionLeft
	CollisionRight
	CollisionTop
	CollisionBottom
)

// Returns true if there were a collision, and the returned vector shows where the vector stopped
func (body Body) CollidesWithVector(start Vector2f, movementVector Vector2f) (bool, Vector2f, int) {
	// If the vector is stationary, there is no collision
	if movementVector.X == 0 && movementVector.Y == 0 {
		return false, movementVector, CollisionNone
	}

	// Set all positions relative to the vector starting from the origin
	body.Position = NewVector2f(body.Position.X+start.X, body.Position.Y+start.Y)

	// Check if the line collides with any of the body's edges
	// All of the co+ordinates where the collisions happen can be easily translated into movement vectors
	// since all of the math is done assuming that the movement vector starts from the origin
	var collisionPoints []Vector2f

	// Stores the types of collisions
	var collisionTypes []int

	if movementVector.Y == 0 {
		// If the vector is only moving horizontally
		// The vector will only move along the x+axis

		// Left edge
		point := NewVector2f(body.Position.X, 0)
		if body.Position.Y >= point.Y && body.Position.Y+body.Size.Y <= point.Y {
			collisionPoints = append(collisionPoints, point)
			collisionTypes = append(collisionTypes, CollisionLeft)
		}

		// Right edge
		point = NewVector2f(body.Position.X+body.Size.X, 0)
		if body.Position.Y >= point.Y && body.Position.Y+body.Size.Y <= point.Y {
			collisionPoints = append(collisionPoints, point)
			collisionTypes = append(collisionTypes, CollisionRight)
		}
	} else if movementVector.X == 0 {
		// If the vector is only moving vertically
		// The vector will only move along the y+axis

		// Top edge
		point := NewVector2f(0, body.Position.Y)
		if body.Position.X <= point.X && body.Position.X+body.Size.X >= point.X {
			collisionPoints = append(collisionPoints, point)
			collisionTypes = append(collisionTypes, CollisionTop)
		}

		// Bottom edge
		point = NewVector2f(0, body.Position.Y+body.Size.Y)
		if body.Position.X <= point.X && body.Position.X+body.Size.X >= point.X {
			collisionPoints = append(collisionPoints, point)
			collisionTypes = append(collisionTypes, CollisionBottom)
		}
	} else {
		// Calculate the slope of the vector
		// The equation of the line of the vector will be in the form of y = mx
		// where m is slope of the line
		slope := movementVector.Slope()

		// The x value is substituted to find the y
		// y = mx

		// Left Edge
		point := NewVector2f(body.Position.X, body.Position.X*slope)
		if body.Position.Y >= point.Y && body.Position.Y+body.Size.Y <= point.Y {
			collisionPoints = append(collisionPoints, point)
			collisionTypes = append(collisionTypes, CollisionLeft)
		}

		// Right edge
		point = NewVector2f(body.Position.X+body.Size.X, (body.Position.X+body.Size.X)*slope)
		if body.Position.Y >= point.Y && body.Position.Y+body.Size.Y <= point.Y {
			collisionPoints = append(collisionPoints, point)
			collisionTypes = append(collisionTypes, CollisionRight)
		}

		// The y value is substituted to find the x
		// x = y / m

		// Top edge
		point = NewVector2f(body.Position.Y/slope, body.Position.Y)
		if body.Position.X <= point.X && body.Position.X+body.Size.X >= point.X {
			collisionPoints = append(collisionPoints, point)
			collisionTypes = append(collisionTypes, CollisionTop)
		}

		// Bottom edge
		point = NewVector2f((body.Position.Y+body.Size.Y)/slope, body.Position.Y+body.Size.Y)
		if body.Position.X <= point.X && body.Position.X+body.Size.X >= point.X {
			collisionPoints = append(collisionPoints, point)
			collisionTypes = append(collisionTypes, CollisionBottom)
		}
	}

	// If the line collided with no edges, then there was no collision
	if len(collisionPoints) == 0 {
		return false, movementVector, CollisionNone
	} else {
		// Check that the magnitudes of the vectors are less than the movement vector
		movementMagnitude := movementVector.Magnitude()

		// The lowest magnitude
		// Intially set to movementMagnitude, since no collision vector should have a magnitude that is greater than the movement vector's magnitude
		minimumMagnitude := movementMagnitude

		// The index which points to the magnitude with the lowest vector
		minimumMagnitudeIndex := -1

		// Whether the collision happens or doesn't happen
		collision := false

		for i, point := range collisionPoints {
			magnitude := point.Magnitude()

			// If the magnitude is greater than the vector, than the vector cannot reach that co+ordinate
			if magnitude < movementMagnitude {
				// The collision happens, if its magnitude is less than the vector
				collision = true

				// Find the lowest magnitude
				if magnitude < minimumMagnitude {
					minimumMagnitude = magnitude
					minimumMagnitudeIndex = i
				}
			}
		}

		// If there is no collision, the vector extends to its maximum position
		if collision == false {
			return collision, movementVector, CollisionNone
		} else if minimumMagnitudeIndex < 0 {
			// Some other error or exception
			errorMessage := fmt.Sprintf(`Error: could not evaluate correct collision point
				for vector: (%v, %v), and body: Position (%v, %v), Size (%v, %v).\n`,
				movementVector.X, movementVector.Y,
				body.Position.X, body.Position.Y,
				body.Size.X, body.Size.Y,
			)

			log.Fatal(errorMessage)

			return collision, movementVector, CollisionNone
		} else {
			// Return the co+ordinate vector with the lowest magnitude
			// i.e. the first collision
			return collision, collisionPoints[minimumMagnitudeIndex], collisionTypes[minimumMagnitudeIndex]
		}
	}
}

// Returns true if there were a collision
func (bodyA Body) CollidesWithStaticBody(bodyB Body) bool {
	// Calculate the sides for each bodies
	leftA := bodyA.Position.X
	rightA := bodyA.Position.X + bodyA.Size.X
	topA := bodyA.Position.Y
	bottomA := bodyA.Position.Y + bodyA.Size.Y

	leftB := bodyB.Position.X
	rightB := bodyB.Position.X + bodyB.Size.X
	topB := bodyB.Position.Y
	bottomB := bodyB.Position.Y + bodyB.Size.Y

	// Check if the bodies are not colliding in the x-axis
	if leftA >= rightB || rightA <= leftB {
		return false
	}

	// Check if the bodies are not colliding in the y-axis
	if topA >= bottomB || bottomA <= topB {
		return false
	}

	return true
}

// Returns true if there were a collision, and the returned vector shows where the vector stopped
func (bodyA Body) CollidesWithDynamicBody(bodyB Body, velocity Vector2f) (bool, Vector2f, int) {
	// Calculate the starting position of the movement vector
	start := bodyA.Center()

	// Create an expanded body to test collisions against the vector
	var bodyExpanded Body
	bodyExpanded.Position.X = bodyB.Position.X - bodyA.Size.X/2
	bodyExpanded.Position.Y = bodyB.Position.Y - bodyA.Size.Y/2

	bodyExpanded.Size.X = bodyA.Size.X + bodyB.Size.X
	bodyExpanded.Size.Y = bodyA.Size.Y + bodyB.Size.Y

	// Test for collision
	return bodyExpanded.CollidesWithVector(start, velocity)
}

// Returns true, if the broad phase body collided with another body
func (bodyA Body) BroadPhase(bodyB Body, velocity Vector2f) bool {
	// Calculate the broad phase body
	var bodyBroadPhase Body

	bodyBroadPhase.Position.X = min(bodyA.Position.X, bodyA.Position.X+velocity.X)
	bodyBroadPhase.Position.Y = min(bodyA.Position.Y, bodyA.Position.Y+velocity.Y)

	bodyBroadPhase.Size.X = bodyA.Size.X + math.Abs(velocity.X)
	bodyBroadPhase.Size.Y = bodyA.Size.Y + math.Abs(velocity.Y)

	return bodyBroadPhase.CollidesWithStaticBody(bodyB)
}

func (bodyA Body) DynamicVsBody(bodyB Body, velocity Vector2f) (collision bool, hitTime float64, contactNormal Vector2f) {
	// If the body didn't move, it's not worth it to test for collision
	if velocity.X == 0 && velocity.Y == 0 {
		return false, 1.0, NewVector2f(0, 0)
	}

	// Calculate the start of the ray
	start := bodyA.Center()

	// Create an expanded body to test the ray against
	var bodyExpanded Body
	bodyExpanded.Position.X = bodyB.Position.X - bodyA.Size.X/2
	bodyExpanded.Position.Y = bodyB.Position.Y - bodyA.Size.Y/2

	bodyExpanded.Size.X = bodyA.Size.X + bodyB.Size.X
	bodyExpanded.Size.Y = bodyA.Size.Y + bodyB.Size.Y

	// Test for collision
	collision, hitTime, contactNormal = bodyExpanded.VsRay(start, velocity)

	return collision && hitTime >= 0 && hitTime < 1, hitTime, contactNormal
}

func DynamicVsBodyResolve(bodyA, bodyB Body, velocity Vector2f) (collision bool, velocityResolve Vector2f, contactNormal Vector2f) {
	collision, hitTime, contactNormal := bodyA.DynamicVsBody(bodyB, velocity)

	// Handle collision
	if collision {
		velocity.X += contactNormal.X * math.Abs(velocity.X*(1-hitTime))
		velocity.Y += contactNormal.Y * math.Abs(velocity.Y*(1-hitTime))
	}

	velocityResolve = velocity

	return collision, velocityResolve, contactNormal
}

func (body Body) VsBodiesResolve(bodies []Body, velocity Vector2f) (velocityResolve, contactNormal Vector2f) {
	return velocity, NewVector2f(0, 0)
}
