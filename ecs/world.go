package ecs

import (
	"sort"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

// Contains entity count information,
// which entity contains which components,
// and component storage, stored as slices
type World struct {
	// Gravity
	Gravity    float32
	MaxGravity float32

	// Entity count
	Size int

	// Component storage
	ComponentPool map[string]any
}

// Entity exist component
type Alive bool

// Create a new world and its entities' components
func NewWorld() World {
	world := World{}

	// Gravity
	world.Gravity = 0.3
	world.MaxGravity = 5

	// ECS
	world.ComponentPool = make(map[string]any)

	// Entity exists
	RegisterComponent[Alive](&world)

	// Register components
	RegisterComponent[gfx.Sprite](&world)
	RegisterComponent[physics.Body](&world)
	RegisterComponent[physics.Force](&world)

	// Entity traits
	RegisterComponent[physics.Jump](&world)

	// Tags
	RegisterComponent[PlayerTag](&world)
	RegisterComponent[EnemyTag](&world)
	RegisterComponent[TileTag](&world)

	// Load maps
	world.LoadMap("assets/maps/map0.json")

	// Create the enemies
	world.NewEnemy()

	// Create the player
	world.NewPlayer()

	return world
}

func (world *World) IsEntityAlive(id int) bool {
	componentSet := *GetComponentSet[Alive](world)

	_, ok := componentSet.Get(id)

	return ok
}

func (world *World) NewEntity() int {
	for id := range world.Size {
		// If the entity is not alive, assign the new entity id
		if !world.IsEntityAlive(id) {
			// Check the entity is now alive
			componentSet := GetComponentSet[Alive](world)
			componentSet.Add(id, true)

			return id
		}
	}

	// If every entity that currently exists is alive, add a new entity position
	id := world.Size

	// Check the entity is now alive
	componentSet := GetComponentSet[Alive](world)
	componentSet.Add(id, true)

	// Increase the number of entities
	world.Size++

	return id
}

func (world *World) DeleteEntity(id int) {
	componentSet := GetComponentSet[Alive](world)
	componentSet.Remove(id)
}

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
		collision := physics.BodyBroadPhase(*body, *tileBody, force.Velocity)

		if !collision {
			continue
		}

		// Check for collision
		collision, _, _ = physics.BodyDynamicVsBody(*body, *tileBody, force.Velocity)

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

		collision, velocityResolve, contactNormal := physics.BodyDynamicVsBodyResolve(*body, *tileBody, force.Velocity)

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
		force.Acceleration.Y += world.Gravity

		// Limit the gravity
		force.Acceleration.Y = min(world.MaxGravity, force.Acceleration.Y)

		// If the body is on the ground, set gravity to zero
		if force.Collisions.Down {
			force.Acceleration.Y = min(0, force.Acceleration.Y)
		}

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

func (world *World) UpdateSprite() {
	// Get all the entities which have the sprite component and the body component
	entities := GetEntities2[gfx.Sprite, physics.Body](world)

	for _, id := range entities {
		// Get the components
		sprite := GetComponent[gfx.Sprite](world, id)
		body := GetComponent[physics.Body](world, id)

		// Update the position of the sprite
		sprite.DstPosition = body.Position
	}
}

func (world *World) Update() {
	// Take in input
	world.UpdateInput()

	// Movement
	world.UpdateMovement()

	// Attacking
	world.EntityAttack()

	// Update the physics world
	world.UpdatePhysics()

	// Update the sprite after all the physics is finished
	world.UpdateSprite()
}

func (world *World) Render() {
	entities := GetEntities[gfx.Sprite](world)

	for _, id := range entities {
		sprite := GetComponent[gfx.Sprite](world, id)

		sprite.Render()
	}
}

func (world *World) Destroy() {
	entities := GetEntities[gfx.Sprite](world)

	for _, id := range entities {
		sprite := GetComponent[gfx.Sprite](world, id)

		sprite.Destroy()
	}
}
