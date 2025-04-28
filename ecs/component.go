package ecs

import (
	// "errors"
	// "fmt"
	"reflect"
)

// Get the type of a generic type interface as reflect.Type variable
func getType[T any]() reflect.Type {
	var temp T
	return reflect.TypeOf(temp)
}

// Register a component with the component type
func RegisterComponent[T any](world *World) {
	// Create a slice for the component at the given index
	componentSlice := make([]T, 0)
	world.ComponentPool[getType[T]()] = &componentSlice
}

func GetComponent[T any](world *World) *[]T {
	// Get the address of the component slice
	componentSliceAddress, ok := world.ComponentPool[getType[T]()].(*[]T)
	if !ok {
		panic("Component type not found")
	}

	return componentSliceAddress
}
