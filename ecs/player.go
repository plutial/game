package ecs

import (
	"log"
	
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

type PlayerTag bool

type Jump struct {
	// Jump "buffers" (like Coyote time)
	AirTime 		int
	JumpRegistered 	int

	// Number of jumps available
	Jumps	int
}

func (world *World) NewPlayer() {
	id := world.NewEntity()

	// Assign a player tag to mark the player entity
	AddComponent[PlayerTag](world, id)

	// Add components
	sprite, err := AddComponent[gfx.Sprite](world, id)

	if err != nil {
		log.Fatal(err)
	}

	*sprite = gfx.NewSprite(gfx.NewTexture("assets/res/image.png"))

	// Body
	body, err := AddComponent[physics.Body](world, id)

	if err != nil {
		log.Fatal(err)
	}

	position := rl.NewVector2(0, 0)
	size := rl.NewVector2(16, 16)

	*body = physics.NewBody(position, size)

	// Force
	_, err = AddComponent[physics.Force](world, id)
	
	if err != nil {
		log.Fatal(err)
	}

	// Jump
	_, err = AddComponent[Jump](world, id)

	if err != nil {
		log.Fatal(err)
	}

	// Colllisions
	_, err = AddComponent[physics.Collisions](world, id)

	if err != nil {
		log.Fatal(err)
	}
}
