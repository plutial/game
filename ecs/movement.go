package ecs

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/physics"
)

// Take in input and change it to movement
func (world *World) UpdateMovement() {
	// Get the player id
	playerId := GetEntities[PlayerTag](world)[0]

	force := GetComponent[physics.Force](world, playerId)

	// Horizontal movement
	force.Move(rl.IsKeyDown(rl.KeyA), rl.IsKeyDown(rl.KeyD))
	force.Dash(rl.IsKeyDown(rl.KeyA), rl.IsKeyDown(rl.KeyD), rl.IsKeyPressed(rl.KeySpace))

	// Update jumps
	jump := GetComponent[physics.Jump](world, playerId)

	force.Jump(jump, rl.IsKeyPressed(rl.KeyW))
}

func (world *World) EntityAttack() {
	// Dismiss if the player does not attack
	if !rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		return
	}

	// Get the player id
	playerId := GetEntities[PlayerTag](world)[0]

	playerBody := GetComponent[physics.Body](world, playerId)

	// Center of the player body
	center := playerBody.Center()

	// Get the enemies
	enemies := GetEntities[EnemyTag](world)

	for _, enemyId := range enemies {
		enemyBody := GetComponent[physics.Body](world, enemyId)
		enemyForce := GetComponent[physics.Force](world, enemyId)

		// Raycast an attack if the enemy is in range
		if physics.GetDistance(playerBody.Position, enemyBody.Position) < 80 {
			movement := rl.NewVector2(
				enemyBody.Center().X-playerBody.Center().X,
				enemyBody.Center().Y-playerBody.Center().Y,
			)

			// Check if the ray is blocked by any of the tiles
			blocked := false

			tiles := GetEntities[TileTag](world)

			for _, tileId := range tiles {
				tileBody := GetComponent[physics.Body](world, tileId)

				// Carry out a broad phase to stop handling
				// Minimize expensive physics on absurd tiles that will never collide with
				collision := playerBody.BroadPhase(*tileBody, movement)

				if !collision {
					continue
				}

				// Check for collision
				collision, _, _ = tileBody.VsRay(center, movement)

				if collision {
					blocked = true
					break
				}
			}

			if !blocked {
				if playerBody.Position.X-enemyBody.Position.X > 0 {
					enemyForce.Velocity.X = -30
				} else {
					enemyForce.Velocity.X = 30
				}

				enemyForce.Velocity.Y = -30
				enemyForce.Acceleration.Y = -0.6
			}
		}
	}
}
