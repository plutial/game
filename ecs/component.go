package ecs

import (
	"fmt"
	"log"

	"github.com/plutial/game/util"
)

// Register a component with the component type
func RegisterComponent[T any](manager *Manager) {
	// Create a slice for the component at the given index
	componentSet := util.NewSparseSet[T]()
	manager.ComponentPool[util.GetType[T]()] = &componentSet
}

// Get the address of the component slice
func GetComponentSet[T any](manager *Manager) *util.SparseSet[T] {
	// Get the address of the component slice
	componentSetAddress, ok := manager.ComponentPool[util.GetType[T]()].(*util.SparseSet[T])

	if !ok {
		message := fmt.Sprintf("Component type %v not found", util.GetType[T]())
		panic(message)
	}

	return componentSetAddress
}

// Check if an entity has a component
func HasComponent[T any](manager *Manager, id int) bool {
	componentSet := GetComponentSet[T](manager)

	_, ok := componentSet.Get(id)

	// Return the check
	return ok
}

// Add a component to an entity
func AddComponent[T any](manager *Manager, id int) *T {
	// Get the component slice
	componentSet := GetComponentSet[T](manager)

	// If the entity already has the component, return the address of the component
	if HasComponent[T](manager, id) {
		address, _ := componentSet.GetAddress(id)
		return address
	}

	// Add the entity
	var temp T
	componentSet.Add(id, temp)
	address, _ := componentSet.GetAddress(id)

	// Return the address of the component
	return address
}

// Remove a component from an entity
func RemoveComponent[T any](manager *Manager, id int) {
	componentSet := GetComponentSet[T](manager)
	componentSet.Delete(id)
}

// Get the address of the component
func GetComponent[T any](manager *Manager, id int) *T {
	componentSet := GetComponentSet[T](manager)

	address, ok := componentSet.GetAddress(id)

	// Check if the entity has the component
	if !ok {
		// Send an error message
		message := fmt.Sprintf(
			"Entity %v is either not alive and/or does not have the component %v",
			id, util.GetType[T](),
		)
		log.Fatal(message)
		return nil
	}

	// Return the address of the component
	return address
}

// Returns a slice of entity ids which have component A
func GetEntities[A any](manager *Manager) []int {
	entities := make([]int, 0)

	for id := range manager.Size {
		// Check that entity is alive
		if !manager.IsEntityAlive(id) {
			continue
		}

		// Add to the slice if the entity has all the required components
		if HasComponent[A](manager, id) {
			entities = append(entities, id)
		}
	}

	return entities
}

// Returns a slice of entity ids which have component A and B
func GetEntities2[A, B any](manager *Manager) []int {
	entities := make([]int, 0)

	for id := range manager.Size {
		// Check that entity is alive
		if !manager.IsEntityAlive(id) {
			continue
		}

		// Add to the slice if the entity has all the required components
		if HasComponent[A](manager, id) && HasComponent[B](manager, id) {
			entities = append(entities, id)
		}
	}

	return entities
}
