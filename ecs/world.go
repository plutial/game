package ecs

import (
	"reflect"
)

type World struct {
	// Entity count
	Size uint32

	// Components
	ComponentPool map[reflect.Type]any
}

func NewWorld() World {
	world := World {}
	world.ComponentPool = make(map[reflect.Type]any)

	return world
}
