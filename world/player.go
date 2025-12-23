package world

import (
	"image/color"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

type PlayerTag bool

func NewPlayer(manager *ecs.Manager) {
	id := manager.NewEntity()

	// Assign a player tag to mark the player entity
	ecs.AddComponent[PlayerTag](manager, id)

	// Add components
	sprite := ecs.AddComponent[gfx.Sprite](manager, id)
	*sprite = gfx.NewSprite(gfx.NewTexture("assets/res/image.png"))
	sprite.Image = nil
	sprite.Color = color.RGBA{0, 255, 0, 255}

	// Body
	body := ecs.AddComponent[physics.Body](manager, id)

	position := physics.NewVector2f(16, 16)
	size := physics.NewVector2f(16, 16)

	*body = physics.NewBody(position, size)

	// Force
	force := ecs.AddComponent[physics.Force](manager, id)
	*force = physics.NewForce(physics.NewVector2f(0, 0), physics.NewVector2f(0, 0))

	// Jump
	ecs.AddComponent[physics.Jump](manager, id)
}
