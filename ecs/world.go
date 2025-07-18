package ecs

import (
	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

// Contains entity count information,
// which entity contains which components,
// and component storage, stored as slices
type World struct {
	// Component storage
	ComponentPool map[string]any

	// Entity count
	Size int

	// Gravity
	Gravity    float32
	MaxGravity float32
}

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

func (world *World) Destroy() {
	world.DestroyEntities()
}
