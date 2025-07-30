package ecs

import (
	"reflect"

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

	// Entities to delete
	ToDelete []int
}

// Create a new world and its entities' components
func NewWorld() World {
	world := World{}

	// ECS
	world.ComponentPool = make(map[string]any)

	// Entity exists
	RegisterComponent[Alive](&world)

	// Register components
	world.RegisterComponents()

	// Load maps
	world.LoadMap("assets/maps/map0.json")

	// Create the enemies
	world.NewEnemy()

	// Create the player
	world.NewPlayer()

	return world
}

func (world *World) RegisterComponents() {
	// Sprite for rendering
	RegisterComponent[gfx.Sprite](world)

	// Physics components
	RegisterComponent[physics.Body](world)
	RegisterComponent[physics.Force](world)

	// Entity traits
	RegisterComponent[physics.Jump](world)

	// Tags
	RegisterComponent[PlayerTag](world)
	RegisterComponent[EnemyTag](world)
	RegisterComponent[TileTag](world)
}

func (world *World) Update() {
	// Delete entities which need to be deleted
	// Delete entities at the start of the loop to manages entites more easily
	world.DeleteEntities()

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

func (world *World) DeleteEntities() {
	for _, id := range world.ToDelete {
		// Check that the entity is alive before removing the alive component
		if !world.IsEntityAlive(id) {
			continue
		}

		// Remove all of the entity's components
		for _, value := range world.ComponentPool {
			value := reflect.ValueOf(value)
			in := make([]reflect.Value, 0)
			in = append(in, reflect.ValueOf(id))

			deleteMethod := value.MethodByName("Delete")
			if deleteMethod.IsValid() {
				deleteMethod.Call(in)
			} else {
				panic("Delete method for sparse set not found")
			}
		}
	}

	// Reset the list
	world.ToDelete = make([]int, 0)
}

func (world *World) Destroy() {
	// Unload all textures
	entities := GetEntities[gfx.Sprite](world)

	for _, id := range entities {
		sprite := GetComponent[gfx.Sprite](world, id)

		sprite.Destroy()
	}
}
