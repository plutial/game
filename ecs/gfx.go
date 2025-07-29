package ecs

import (
	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

func (world *World) UpdateSprite() {
	// Get all the entities which have the sprite component and the body component
	entities := GetEntities2[gfx.Sprite, physics.Body](world)

	for _, id := range entities {
		// Get the components
		sprite := GetComponent[gfx.Sprite](world, id)
		body := GetComponent[physics.Body](world, id)

		// Update the position of the sprite
		sprite.Destination.Position = body.Position
	}
}
