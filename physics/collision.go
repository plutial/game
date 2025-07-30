package physics

import (
	"math"
)

func (bodyA *Body) StaticVsBody(bodyB Body) (collision bool) {
	// Get the collisions for each axis
	xCollision := bodyA.Position.X < bodyB.Position.X+bodyB.Size.X ||
		bodyA.Position.X+bodyA.Size.X > bodyB.Position.X

	yCollision := bodyA.Position.Y < bodyB.Position.Y+bodyB.Size.Y ||
		bodyA.Position.Y+bodyA.Size.Y > bodyB.Position.Y

	return xCollision && yCollision
}

func (bodyA *Body) BroadPhase(bodyB Body, velocity Vector2) (collision bool) {
	// Calculate the broad phase body
	var bodyBroadPhase Body
	if velocity.X > 0 {
		bodyBroadPhase.Position.X = bodyA.Position.X
		bodyBroadPhase.Position.Y = bodyA.Position.Y

		bodyBroadPhase.Size.X = bodyA.Size.X + velocity.X
		bodyBroadPhase.Size.Y = bodyA.Size.Y + velocity.Y
	} else {
		bodyBroadPhase.Position.X = bodyA.Position.X + velocity.X
		bodyBroadPhase.Position.Y = bodyA.Position.Y + velocity.Y

		bodyBroadPhase.Size.X = bodyA.Size.X - velocity.X
		bodyBroadPhase.Size.Y = bodyA.Size.Y - velocity.Y
	}

	return bodyBroadPhase.StaticVsBody(bodyB)
}

func (bodyA *Body) DynamicVsBody(bodyB Body, velocity Vector2) (collision bool, hitTime float64, contactNormal Vector2) {
	// If the body didn't move, it's not worth it to test for collision
	if velocity.X == 0 && velocity.Y == 0 {
		return false, 1.0, NewVector2(0, 0)
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

func DynamicVsBodyResolve(bodyA, bodyB Body, velocity Vector2) (collision bool, velocityResolve Vector2, contactNormal Vector2) {
	collision, hitTime, contactNormal := bodyA.DynamicVsBody(bodyB, velocity)

	// Handle collision
	if collision {
		velocity.X += contactNormal.X * math.Abs(velocity.X*(1-hitTime))
		velocity.Y += contactNormal.Y * math.Abs(velocity.Y*(1-hitTime))
	}

	velocityResolve = velocity

	return collision, velocityResolve, contactNormal
}

func (body *Body) VsBodiesResolve(bodies []Body, velocity Vector2) (velocityResolve, contactNormal Vector2) {
	return velocity, NewVector2(0, 0)
}
