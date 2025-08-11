package world

import (
	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

func UpdateSprite(manager *ecs.Manager) {
	// Get all the entities which have the sprite component and the body component
	entities := ecs.GetEntities2[gfx.Sprite, physics.Body](manager)

	for _, id := range entities {
		// Get the components
		sprite := ecs.GetComponent[gfx.Sprite](manager, id)
		body := ecs.GetComponent[physics.Body](manager, id)

		// Update the position of the sprite
		sprite.Destination.Position = body.Position
	}
}
