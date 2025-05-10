package ecs

import (
	"fmt"
)

// Register a component with the component type
func RegisterComponent[T any](world *World) {
	// Create a slice for the component at the given index
	componentSlice := make([]T, 1)
	world.ComponentPool[getType[T]()] = &componentSlice
}

// Get the address of the component slice
func GetComponentSlice[T any](world *World) *[]T {
	// Get the address of the component slice
	componentSliceAddress, ok := world.ComponentPool[getType[T]()].(*[]T)

	if !ok {
		panic("Component type not found")
	}

	return componentSliceAddress
}

// Check if an entity has a component
func HasComponent[T any](world *World, id int) bool {
	hasComponent, ok := world.EntityHasComponent[id]
	
	// If the entity is dead, return false
	if !ok {
		return false
	}

	// Check if the entity has the said component
	_, ok = hasComponent[getType[T]()]

	// Return the check
	return ok
}

// Add a component to an entity
func AddComponent[T any](world *World, id int) (*T, error) {
	// Get the component slice
	componentSlice := GetComponentSlice[T](world)

	// If the entity already has the component, return the address of the component
	if HasComponent[T](world, id) {
		return &(*componentSlice)[id], nil
	}

	// Check if the entity exists
	_, ok := world.EntityHasComponent[id]

	// If the entity is dead, return an error
	if !ok {
		return nil, fmt.Errorf("Entity %v is not alive; therefore, adding the component %v is not possible", id, getType[T]())
	}

	// Mark that the entity has the component
	world.EntityHasComponent[id][getType[T]()] = true

	// If the id is out of range of the slice, then double size of the slice
	if cap(*componentSlice) <= id {
		*componentSlice = append(*componentSlice, make([]T, cap(*componentSlice))...)
	}
	
	// Return the address of the component
	return &(*componentSlice)[id], nil
}

// Get the address of the component
func GetComponent[T any](world *World, id int) (*T, error) {
	// Check if the entity has the component
	if !HasComponent[T](world, id) {
		return nil, fmt.Errorf("Entity %v is either not alive and/or does not have the component %v", id, getType[T]())
	}

	componentSlice := GetComponentSlice[T](world) 

	// Return the address of the component
	return &(*componentSlice)[id], nil
}

func GetEntities[A any](world *World) []int {
	entities := make([]int, 0)

	for id := range world.Size {
		if HasComponent[A](world, id) {
			entities = append(entities, id)
		}
	}
	
	return entities
}

func GetEntities2[A, B any](world *World) []int {
	entities := make([]int, 0)

	for id := range world.Size {
		if HasComponent[A](world, id) && HasComponent[B](world, id) {
			entities = append(entities, id)
		}
	}
	
	return entities
}
