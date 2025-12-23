package world

import (
	"fmt"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/physics"
)

// Update tile physics against a body with force
func UpdateTilePhysics(manager *ecs.Manager, body *physics.Body, force *physics.Force, tiles []int) {
}

// Update all the entites with a body and force
func UpdatePhysics(manager *ecs.Manager) {
	// Get all the entities which have the body component and the force component
	entities := ecs.GetEntities2[physics.Body, physics.Force](manager)

	// Get all tiles
	tiles := ecs.GetEntities[TileTag](manager)

	// Get all tile bodies
	var tileBodies []physics.Body

	for _, id := range tiles {
		// Get the tile body
		body := ecs.GetComponent[physics.Body](manager, id)

		// Add the tile body
		tileBodies = append(tileBodies, *body)
	}

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
		fmt.Println(force.Velocity)
		body.CollidiesWithDynamicBodies(tileBodies, force)

		fmt.Println(force.Velocity)
		fmt.Println(force.Collisions, "\n")

		// Update the body position
		body.Position.X += force.Velocity.X
		body.Position.Y += force.Velocity.Y

		// Reset the velocity after calculation
		force.Velocity.X = 0
		force.Velocity.Y = 0
	}
}
