package ecs

import (
	"image/color"

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
	sprite.Image = nil
	sprite.Color = color.RGBA{0, 255, 0, 255}

	// Body
	body := AddComponent[physics.Body](world, id)

	position := physics.NewVector2(16, 16)
	size := physics.NewVector2(16, 16)

	*body = physics.NewBody(position, size)

	// Force
	force := AddComponent[physics.Force](world, id)
	*force = physics.NewForce(physics.NewVector2(0, 0), physics.NewVector2(0, 0))

	// Jump
	AddComponent[physics.Jump](world, id)
}
