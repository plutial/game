package ecs

import (
	"sort"

	// Game packages
	"github.com/plutial/game/physics"
)

// Update tile physics against a body with force
func (world *World) UpdateTilePhysics(body *physics.Body, force *physics.Force, tiles []int) {
	// Slice to store tiles which have collided
	type TileCollisionData struct {
		TileId   int
		Distance float32
	}

	tileCollisionData := make([]TileCollisionData, 0)

	// Check which tiles could collided with the body
	for _, tileId := range tiles {
		tileBody := GetComponent[physics.Body](world, tileId)

		// Carry out a broad phase to stop handling
		// Minimize expensive physics on absurd tiles that will never collide with
		collision := body.BroadPhase(*tileBody, force.Velocity)

		if !collision {
			continue
		}

		// Check for collision
		collision, _, _ = body.DynamicVsBody(*tileBody, force.Velocity)

		if collision {
			// Get the distance from the body to the tile
			distance := physics.GetDistance(body.Position, tileBody.Position)

			data := TileCollisionData{tileId, distance}

			// Add to the collided tile list
			tileCollisionData = append(tileCollisionData, data)
		}
	}

	// Sort the tiles by which tiles are the closest to the body
	// This is a fix to imitate actual physics (and to handle other cases)
	sort.SliceStable(tileCollisionData, func(a, b int) bool {
		return tileCollisionData[a].Distance < tileCollisionData[b].Distance
	})

	// Reset the collisions
	force.Collisions = physics.Collisions{}

	// Resolve the collisions
	for _, data := range tileCollisionData {
		tileId := data.TileId

		tileBody := GetComponent[physics.Body](world, tileId)

		collision, velocityResolve, contactNormal := physics.DynamicVsBodyResolve(*body, *tileBody, force.Velocity)

		if collision {
			// Update the collision velocity
			force.Velocity = velocityResolve

			// Update the collision direction
			force.Collisions.Update(contactNormal)
		}
	}
}

// Update all the entites with a body and force
func (world *World) UpdatePhysics() {
	// Get all the entities which have the body component and the force component
	entities := GetEntities2[physics.Body, physics.Force](world)

	// Get all tiles
	tiles := GetEntities[TileTag](world)

	for _, id := range entities {
		// Get the components
		body := GetComponent[physics.Body](world, id)
		force := GetComponent[physics.Force](world, id)

		// Apply gravity
		force.UpdateGravity()

		// Update acceleration
		force.Velocity.X += force.Acceleration.X
		force.Velocity.Y += force.Acceleration.Y

		// Handle tile collisions
		// This MUST be handled at the end AFTER acceleration has been applied
		world.UpdateTilePhysics(body, force, tiles)

		// Update the body position
		body.Position.X += force.Velocity.X
		body.Position.Y += force.Velocity.Y

		// Reset the velocity after calculation
		force.Velocity.X = 0
		force.Velocity.Y = 0
	}
}
