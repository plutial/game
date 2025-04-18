package ecs

import (
	"fmt"
	"errors"
	"reflect"
)

func getType[T any]() reflect.Type {
	// Create a temporary slice
	temp := make([]T, 0)
	return reflect.TypeOf(temp)
}

func getComponentSlice[T any](world World) ([]T, error) {
	// Get the component slice
	component, ok := world.ComponentPool[getType[T]()].([]T)

	// Check if the component exists
	if !ok {
		return make([]T, 0, 0), errors.New("Component type not found")
	}

	return component, nil
}

func RegisterComponent[T any](world *World) []T {
	// Create an array of T and add it to the component pool
	world.ComponentPool[getType[T]()] = make([]T, 0)
	return world.ComponentPool[getType[T]()].([]T)
}

func AddComponent[T any](world *World, entity uint32) {
	component, err := getComponentSlice[T](*world)
	
	if err != nil {
		fmt.Println(err)
	}

	// If the entity goes past the avaible slice, expand the slice
	if cap(component) > int(entity) {
		buffer := make([]T, cap(component) - int(entity), cap(component) - int(entity))
		component = append(component, buffer...)
	}

	// Update changes
	world.ComponentPool[getType[T]()] = component
}

func GetComponent[T any](world *World, entity uint32) *T {
	// Return the address of the component
	component := (world.ComponentPool[getType[T]()].([]T))
	return &component
}
