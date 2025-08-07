package ecs

import (
	"image/color"

	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/input"
	"github.com/plutial/game/physics"
)

// Take in input and change it to movement
func (world *World) UpdateMovement() {
	// Get the player id
	playerId := GetEntities[PlayerTag](world)[0]

	force := GetComponent[physics.Force](world, playerId)

	// Horizontal movement
	force.Move(input.IsKeyDown(input.KeyA), input.IsKeyDown(input.KeyD))
	force.Dash(input.IsKeyDown(input.KeyA), input.IsKeyDown(input.KeyD), input.IsKeyPressed(input.KeySpace))

	// Update jumps
	jump := GetComponent[physics.Jump](world, playerId)

	force.Jump(jump, input.IsKeyPressed(input.KeyW))
}

func (world *World) EntityAttack() {
	// Dismiss if the player does not attack
	if !input.IsMouseButtonPressed(input.MouseButtonLeft) {
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
		if playerBody.Position.Distance(enemyBody.Position) < 80 {
			movement := physics.NewVector2(
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

type ProjectileTag bool

func (world *World) EntityCharge() {
	if input.IsMouseButtonPressed(input.MouseButtonLeft) {
		// Create a new charge projectile
		id := world.NewEntity()

		// Add a projectile tag so that it explodes when it collides with something
		AddComponent[ProjectileTag](world, id)

		// Get the player position
		playerId := GetEntities[PlayerTag](world)[0]

		playerBody := GetComponent[physics.Body](world, playerId)

		body := AddComponent[physics.Body](world, id)
		*body = physics.NewBody(
			playerBody.Center(),
			physics.NewVector2(8, 8),
		)

		// Make the projectile go in the position of the mouse
		force := AddComponent[physics.Force](world, id)

		projectileSpeed := 1.5
		center := body.Center()
		center.X += force.Acceleration.X
		center.Y += force.Acceleration.Y
		scaleFactor := projectileSpeed / center.Distance(input.MousePosition())

		force.Acceleration.X = input.MousePosition().X - center.X
		force.Acceleration.X *= scaleFactor
		force.Acceleration.Y = input.MousePosition().Y - center.Y
		force.Acceleration.Y *= scaleFactor

		// Add a sprite
		sprite := AddComponent[gfx.Sprite](world, id)
		*sprite = gfx.NewSprite(gfx.NewTexture("assets/res/image.png"))
		sprite.Image = nil
		sprite.Color = color.RGBA{255, 255, 255, 255}
	}

	// Get projectiles
	projectiles := GetEntities[ProjectileTag](world)

	for _, id := range projectiles {
		// Check to see if the projectile collided with anything
		force := GetComponent[physics.Force](world, id)

		// Boost the entities with an explosion
		if force.Collisions.Collided() {
			// Projectile body
			body := GetComponent[physics.Body](world, id)

			// Entities which have a body and a velocity
			entities := GetEntities2[physics.Body, physics.Force](world)

			for _, id := range entities {
				// The explosion does not affect projectiles
				if HasComponent[ProjectileTag](world, id) {
					continue
				}

				entityBody := GetComponent[physics.Body](world, id)
				entityForce := GetComponent[physics.Force](world, id)

				explosion := physics.NewVector2(5, 6.5)
				// If the entity is in range
				center := body.Center()
				if center.Distance(entityBody.Center()) < 32 {
					if body.Center().X-entityBody.Center().X > 0 {
						entityForce.Acceleration.X -= explosion.X
					} else {
						entityForce.Acceleration.X += explosion.X
					}

					if body.Center().Y-entityBody.Center().Y > 0 {
						entityForce.Acceleration.Y -= explosion.Y
					} else {
						entityForce.Acceleration.Y += explosion.Y
					}
				}
			}

			// Remove the projectile
			world.DeleteEntity(id)
		}
	}
}
