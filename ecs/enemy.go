package ecs

import (
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
	AddComponent[EnemyTag](world, id)

	// Add components
	sprite := AddComponent[gfx.Sprite](world, id)

	*sprite = gfx.NewSprite(gfx.NewTexture("assets/res/image.png"))

	// Body
	body := AddComponent[physics.Body](world, id)

	position := rl.NewVector2(50, 0)
	size := rl.NewVector2(16, 16)

	*body = physics.NewBody(position, size)

	// Force
	force := AddComponent[physics.Force](world, id)

	*force = physics.NewForce(rl.NewVector2(0, 0), rl.NewVector2(0, 0))

	// Jump
	AddComponent[physics.Jump](world, id)
}
