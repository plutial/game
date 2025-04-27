package ecs

import (
	"reflect"
)

type World struct {
	// Entity count
	Size uint32

	// Component index
	ComponentIndex map[reflect.Type]int

	// Component storage
	ComponentPool []any
}

func NewWorld() World {
	world := World {}
	world.ComponentIndex = make(map[reflect.Type]int)
	world.ComponentPool = make([]any, 0)

	return world
}
