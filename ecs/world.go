package ecs

import (
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
	// Component storage
	ComponentPool map[string]any

	// Entity count
	Size int

	// Delta time
	DeltaTime    float32
	CurrentTime  float32
	PreviousTime float32

	// Gravity
	Gravity    float32
	MaxGravity float32
}

// Create a new world and its entities' components
func NewWorld() World {
	world := World{}

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

	// Gravity
	world.Gravity = 0.3
	world.MaxGravity = 5

	return world
}

func (world *World) Update() {
	// Take in input and change it to movement
	world.UpdateMovement()

	// Attacking
	world.EntityAttack()

	// Update the physics world
	world.UpdatePhysics()

	// Update the sprite after all the physics calculations have finished
	world.UpdateSprite()
}

func (world *World) Render() {
	// Begin rendering
	rl.BeginDrawing()

	// Clear renderer with a white background
	rl.ClearBackground(rl.RayWhite)

	// Get the entities which have the sprite component
	entities := GetEntities[gfx.Sprite](world)

	// For each entity, render its sprite
	for _, id := range entities {
		sprite := GetComponent[gfx.Sprite](world, id)

		sprite.Render()
	}

	// End renderering and swap buffers
	rl.EndDrawing()
}

func (world *World) Destroy() {
	// Unload all textures
	entities := GetEntities[gfx.Sprite](world)

	for _, id := range entities {
		sprite := GetComponent[gfx.Sprite](world, id)

		sprite.Destroy()
	}
}
