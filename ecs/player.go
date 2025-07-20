package ecs

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

type PlayerTag bool

func (world *World) NewPlayer() {
	id := world.NewEntity()

	// Assign a player tag to mark the player entity
	AddComponent[PlayerTag](world, id)

	// Add components
	sprite := AddComponent[gfx.Sprite](world, id)
	*sprite = gfx.NewSprite(gfx.NewTexture("assets/res/image.png"))
	sprite.Texture.ID = 0
	sprite.Color = rl.NewColor(0, 255, 0, 255)

	// Body
	body := AddComponent[physics.Body](world, id)

	position := rl.NewVector2(0, 0)
	size := rl.NewVector2(16, 16)

	*body = physics.NewBody(position, size)

	// Force
	force := AddComponent[physics.Force](world, id)
	*force = physics.NewForce(rl.NewVector2(0, 0), rl.NewVector2(0, 0))

	// Jump
	AddComponent[physics.Jump](world, id)
}
