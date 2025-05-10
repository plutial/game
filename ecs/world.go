package ecs

import (
	"reflect"
	"log"

	// Game packages
	"github.com/plutial/game/physics"
	"github.com/plutial/game/gfx"
)

// Get the type of a generic type interface as reflect.Type variable
func getType[T any]() reflect.Type {
	var temp T
	return reflect.TypeOf(temp)
}

// Contains entity count information,
// which entity contains which components,
// and component storage, stored as slices
type World struct {
	// Entity count
	Size int

	// Check if an entity has a component
	EntityHasComponent map[int]map[reflect.Type]bool

	// Component storage
	ComponentPool map[reflect.Type]any
}

// Create a new world
func NewWorld() World {
	world := World {}
	world.EntityHasComponent = make(map[int]map[reflect.Type]bool)
	world.ComponentPool = make(map[reflect.Type]any)

	return world
}

func (world *World) NewEntity() int {
	for id := range world.Size {
		_, ok := world.EntityHasComponent[id]
		
		// If the entity is not alive, assign the new entity id
		if !ok {
			// Check the entity is now alive
			world.EntityHasComponent[id] = make(map[reflect.Type]bool)
			return id 
		}
	}

	// If every entity that currently exists is alive, add a new entity position
	id := world.Size
			
	// Check the entity is now alive
	world.EntityHasComponent[id] = make(map[reflect.Type]bool)

	// Increase the number of entities
	world.Size++ 

	return id
}

func (world *World) UpdatePhysics() {
	// Get all the entities which have the body component and the force component
	entities := GetEntities2[physics.Body, physics.Force](world)

	for id := range entities {
		// Get the components
		body, err := GetComponent[physics.Body](world, id)

		if err != nil {
			log.Fatal(err)
		}

		force, err := GetComponent[physics.Force](world, id)

		if err != nil {
			log.Fatal(err)
		}

		// Update the body position
		body.Position.X += force.Velocity.X
		body.Position.Y += force.Velocity.Y
	}
}

func (world *World) UpdateSprite() {
	// Get all the entities which have the sprite component and the body component
	entities := GetEntities2[gfx.Sprite, physics.Body](world)

	for id := range entities {
		// Get the components
		sprite, err := GetComponent[gfx.Sprite](world, id)

		if err != nil {
			log.Fatal(err)
		}

		body, err := GetComponent[physics.Body](world, id)
	
		if err != nil {
			log.Fatal(err)
		}

		// Update the position of the sprite
		*&sprite.DstPosition = *&body.Position
	}
}

func (world *World) Update() {
	world.UpdateSprite()
	world.UpdatePhysics()
}

func (world *World) Render() {
	entities := GetEntities[gfx.Sprite](world)
	
	for id := range entities {
		sprite, err := GetComponent[gfx.Sprite](world, id)

		if err != nil {
			log.Fatal(err)
		}

		sprite.Render()
	}
}
