package ecs

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/physics"
)

func (world *World) UpdateInput() {
	// Get the player id
	playerId := GetEntities[PlayerTag](world)[0]

	force := GetComponent[physics.Force](world, playerId)

	// Horizontal movement
	physics.BodyMove(force, rl.IsKeyDown(rl.KeyA), rl.IsKeyDown(rl.KeyD))
	physics.BodyDash(force, rl.IsKeyDown(rl.KeyA), rl.IsKeyDown(rl.KeyD), rl.IsKeyPressed(rl.KeySpace))
}

func (world *World) UpdateJump() {
	playerId := GetEntities[PlayerTag](world)[0]

	force := GetComponent[physics.Force](world, playerId)
	jump := GetComponent[physics.Jump](world, playerId)

	physics.BodyJump(force, jump, rl.IsKeyPressed(rl.KeyW))
}

func (world *World) UpdateMovement() {
	world.UpdateJump()
}

func (world *World) EntityAttack() {
	// Dismiss if the player does not attack
	if !rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		return
	}

	// Get the player id
	playerId := GetEntities[PlayerTag](world)[0]

	playerBody := GetComponent[physics.Body](world, playerId)

	// Get the enemies
	enemies := GetEntities[EnemyTag](world)

	for _, enemyId := range enemies {
		enemyBody := GetComponent[physics.Body](world, enemyId)
		enemyForce := GetComponent[physics.Force](world, enemyId)

		if physics.GetDistance(playerBody.Position, enemyBody.Position) < 80 {
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
