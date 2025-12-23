package world

import (
	"sort"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/physics"
)

// Update tile physics against a body with force
func UpdateTilePhysics(manager *ecs.Manager, body *physics.Body, force *physics.Force, tiles []int) {
	// Slice to store tiles which have collided
	type TileCollisionData struct {
		// ID of the tile
		TileId int

		// Magnitude of the resultant vector
		Magnitude float64
	}

	tileCollisionData := make([]TileCollisionData, 0)

	// Check which tiles could collided with the body
	for _, tileId := range tiles {
		tileBody := ecs.GetComponent[physics.Body](manager, tileId)

		// Carry out a broad phase to stop handling
		// Minimize expensive physics on absurd tiles that will never collide with
		collision := body.BroadPhase(*tileBody, force.Velocity)

		if collision {
			// Check for collision
			collision, resolvedVelocity, _ := body.CollidesWithDynamicBody(*tileBody, force.Velocity)

			if collision {
				// Get the distance from the body to the tile
				data := TileCollisionData{tileId, resolvedVelocity.Magnitude()}

				// Add to the collided tile list
				tileCollisionData = append(tileCollisionData, data)
			}
		}
	}

	// Sort the tiles by which tiles are the closest to the body
	// This is a fix to imitate actual physics (and to handle other cases)
	sort.SliceStable(tileCollisionData, func(a, b int) bool {
		return tileCollisionData[a].Magnitude < tileCollisionData[b].Magnitude
	})

	// Reset the collisions
	force.Collisions = physics.Collisions{}

	// Resolve the collisions
	for _, data := range tileCollisionData {
		tileId := data.TileId

		tileBody := ecs.GetComponent[physics.Body](manager, tileId)

		collision, velocityResolve, collisionType := body.CollidesWithDynamicBody(*tileBody, force.Velocity)

		if collision {
			// Update the collision velocity
			force.Velocity = velocityResolve

			// Update the collision direction
			switch collisionType {
			case physics.CollisionLeft:
				force.Collisions.Update(physics.NewVector2f(1, 0))
			case physics.CollisionRight:
				force.Collisions.Update(physics.NewVector2f(-1, 0))
			case physics.CollisionTop:
				force.Collisions.Update(physics.NewVector2f(0, 1))
			case physics.CollisionBottom:
				force.Collisions.Update(physics.NewVector2f(0, -1))
			default:
				force.Collisions.Update(physics.NewVector2f(0, 0))
			}
		}
	}
}

// Update all the entites with a body and force
func UpdatePhysics(manager *ecs.Manager) {
	// Get all the entities which have the body component and the force component
	entities := ecs.GetEntities2[physics.Body, physics.Force](manager)

	// Get all tiles
	tiles := ecs.GetEntities[TileTag](manager)

	for _, id := range entities {
		// Get the components
		body := ecs.GetComponent[physics.Body](manager, id)
		force := ecs.GetComponent[physics.Force](manager, id)

		// Apply gravity
		// Projectiles aren't affected
		if !ecs.HasComponent[ProjectileTag](manager, id) {
			force.UpdateGravity()
		}

		// Apply friction
		if !ecs.HasComponent[ProjectileTag](manager, id) {
			force.Friction()
		}

		// Update acceleration
		force.Velocity.X += force.Acceleration.X
		force.Velocity.Y += force.Acceleration.Y

		// Handle tile collisions
		// This MUST be handled at the end AFTER acceleration has been applied
		UpdateTilePhysics(manager, body, force, tiles)

		// Update the body position
		body.Position.X += force.Velocity.X
		body.Position.Y += force.Velocity.Y

		// Reset the velocity after calculation
		force.Velocity.X = 0
		force.Velocity.Y = 0
	}
}
