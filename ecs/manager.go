package ecs

import (
	"reflect"

	// Game packages
	"github.com/plutial/game/gfx"
)

// ECS stands for the "Entity Component System".
// The entity manager contains entity count information,
// which entity contains which components,
// and component storage, stored as slices.
type Manager struct {
	// Component storage
	ComponentPool map[string]any

	// Entity count
	Size int

	// Entities to delete
	ToDelete []int
}

// Create new Manager and its entities' components
func NewManager() Manager {
	manager := Manager{}

	// Manager
	manager.ComponentPool = make(map[string]any)

	// Entity exists
	RegisterComponent[Alive](&manager)

	// Register components
	manager.RegisterComponents()

	return manager
}

func (manager *Manager) RegisterComponents() {
}

func (manager *Manager) Update() {
	// Delete entities which need to be deleted
	// Delete entities at the start of the loop to manages entites more easily
	manager.DeleteEntities()
}

func (manager *Manager) Render() {
	// Get the entities which have the sprite component
	entities := GetEntities[gfx.Sprite](manager)

	// For each entity, render its sprite
	for _, id := range entities {
		sprite := GetComponent[gfx.Sprite](manager, id)

		sprite.Render()
	}
}

func (manager *Manager) DeleteEntities() {
	for _, id := range manager.ToDelete {
		// Check that the entity is alive before removing the alive component
		if !manager.IsEntityAlive(id) {
			continue
		}

		// Remove all of the entity's components
		for _, value := range manager.ComponentPool {
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
	manager.ToDelete = make([]int, 0)
}
