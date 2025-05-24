package ecs

import (
	"log"

	// Game packages
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

type PlayerTag bool

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
	AddComponent[physics.Force](world, id)
}
