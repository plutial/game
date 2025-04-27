package ecs

import (
	// "errors"
	"fmt"
	"reflect"
)

// Get the type of a generic type interface as reflect.Type variable
func getType[T any]() reflect.Type {
	var temp T
	return reflect.TypeOf(temp)
}

// Register a component with the component type
func RegisterComponent[T any](world *World) {
	// Assign the component index to the component
	componentIndex := len(world.ComponentIndex)
	world.ComponentIndex[getType[T]()] = componentIndex

	// Create a slice for the component at the given index
	componentSlice := make([]T, 0)
	world.ComponentPool = append(world.ComponentPool, &componentSlice)

	// Debug
	fmt.Println(world.ComponentPool)
	fmt.Println(world.ComponentIndex)
	fmt.Println(len(world.ComponentPool))
}

func GetComponent[T any](world *World) *[]T {
	// Get the component index
	componentIndex := world.ComponentIndex[getType[T]()]

	// Return the address of the slice
	return world.ComponentPool[componentIndex].(*[]T)
}
