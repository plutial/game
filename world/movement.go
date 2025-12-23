package world

import (
	"image/color"
	"math"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/input"
	"github.com/plutial/game/physics"
)

// Take in input and change it to movement
func UpdateMovement(manager *ecs.Manager) {
	// Get the player id
	playerId := ecs.GetEntities[PlayerTag](manager)[0]

	force := ecs.GetComponent[physics.Force](manager, playerId)

	// Horizontal movement
	force.Move(input.IsKeyDown(input.KeyA), input.IsKeyDown(input.KeyD))
	force.Dash(input.IsKeyDown(input.KeyA), input.IsKeyDown(input.KeyD), input.IsKeyPressed(input.KeySpace))

	// Update jumps
	jump := ecs.GetComponent[physics.Jump](manager, playerId)

	force.Jump(jump, input.IsKeyPressed(input.KeyW))
}

func EntityAttack(manager *ecs.Manager) {
	// Dismiss if the player does not attack
	if !input.IsMouseButtonPressed(input.MouseButtonLeft) {
		return
	}

	// Get the player id
	playerId := ecs.GetEntities[PlayerTag](manager)[0]

	playerBody := ecs.GetComponent[physics.Body](manager, playerId)

	// Center of the player body
	center := playerBody.Center()

	// Get the enemies
	enemies := ecs.GetEntities[EnemyTag](manager)

	for _, enemyId := range enemies {
		enemyBody := ecs.GetComponent[physics.Body](manager, enemyId)
		enemyForce := ecs.GetComponent[physics.Force](manager, enemyId)

		// Raycast an attack if the enemy is in range
		if playerBody.Position.Distance(enemyBody.Position) < 80 {
			movement := physics.NewVector2f(
				enemyBody.Center().X-playerBody.Center().X,
				enemyBody.Center().Y-playerBody.Center().Y,
			)

			// Check if the ray is blocked by any of the tiles
			blocked := false

			tiles := ecs.GetEntities[TileTag](manager)

			for _, tileId := range tiles {
				tileBody := ecs.GetComponent[physics.Body](manager, tileId)

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

func EntityCharge(manager *ecs.Manager) {
	if input.IsMouseButtonPressed(input.MouseButtonLeft) {
		// Create a new charge projectile
		id := manager.NewEntity()

		// Add a projectile tag so that it explodes when it collides with something
		ecs.AddComponent[ProjectileTag](manager, id)

		// Get the player position
		playerId := ecs.GetEntities[PlayerTag](manager)[0]

		playerBody := ecs.GetComponent[physics.Body](manager, playerId)

		body := ecs.AddComponent[physics.Body](manager, id)
		*body = physics.NewBody(
			playerBody.Center(),
			physics.NewVector2f(8, 8),
		)

		// Make the projectile go in the position of the mouse
		force := ecs.AddComponent[physics.Force](manager, id)

		projectileSpeed := 1.5
		center := body.Center()
		center.X -= body.Size.X / 2
		center.Y -= body.Size.Y / 2
		center.X += force.Acceleration.X
		center.Y += force.Acceleration.Y
		scaleFactor := projectileSpeed / center.Distance(input.MousePosition())

		force.Acceleration.X = input.MousePosition().X - center.X
		force.Acceleration.X *= scaleFactor
		force.Acceleration.Y = input.MousePosition().Y - center.Y
		force.Acceleration.Y *= scaleFactor

		// Add a sprite
		sprite := ecs.AddComponent[gfx.Sprite](manager, id)
		*sprite = gfx.NewSprite(gfx.NewTexture("assets/res/image.png"))
		sprite.Image = nil
		sprite.Color = color.RGBA{255, 255, 255, 255}
		sprite.Destination.Size = physics.NewVector2f(8, 8)

		// The rotation
		// The vertical acceleration is equal to the length opposite side
		// The horizontal acceleration is to the length adjacent side
		sprite.Rotation = math.Atan(force.Acceleration.Y / force.Acceleration.X)
	}

	// Get projectiles
	projectiles := ecs.GetEntities[ProjectileTag](manager)

	for _, id := range projectiles {
		// Check to see if the projectile collided with anything
		force := ecs.GetComponent[physics.Force](manager, id)

		// Boost the entities with an explosion
		if force.Collisions.Collided() {
			// Projectile body
			body := ecs.GetComponent[physics.Body](manager, id)

			// Entities which have a body and a velocity
			entities := ecs.GetEntities2[physics.Body, physics.Force](manager)

			for _, id := range entities {
				// The explosion does not affect projectiles
				if ecs.HasComponent[ProjectileTag](manager, id) {
					continue
				}

				entityBody := ecs.GetComponent[physics.Body](manager, id)
				entityForce := ecs.GetComponent[physics.Force](manager, id)

				explosion := physics.NewVector2f(5, 6.5)
				// If the entity is in range
				center := body.Center()
				center.X -= body.Size.X / 2
				center.Y -= body.Size.Y / 2
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
			manager.DeleteEntity(id)
		}
	}
}
