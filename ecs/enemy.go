package ecs

import (
	"log"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

type EnemyTag bool

func (world *World) NewEnemy() {
	id := world.NewEntity()

	// Assign a player tag to mark the player entity
	_, err := AddComponent[EnemyTag](world, id)
	if err != nil {
		log.Fatal(err)
	}

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

	position := rl.NewVector2(50, 0)
	size := rl.NewVector2(16, 16)

	*body = physics.NewBody(position, size)

	// Force
	force, err := AddComponent[physics.Force](world, id)
	if err != nil {
		log.Fatal(err)
	}

	*force = physics.NewForce(rl.NewVector2(0, 0), rl.NewVector2(0, 0))

	// Jump
	_, err = AddComponent[physics.Jump](world, id)
	if err != nil {
		log.Fatal(err)
	}
}
