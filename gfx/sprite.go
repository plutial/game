package gfx

import (
	// Game packages
	"github.com/plutial/game/physics"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	// The texture itself
	Texture rl.Texture2D

	// The color the sprite will render with if the texture is nil
	Color rl.Color

	// The source co-ordinates and size
	Source physics.Body

	// The destination co-ordinates and size
	Destination physics.Body
}

func NewSprite(texture rl.Texture2D) Sprite {
	var sprite Sprite

	// Load the texture
	sprite.Texture = texture

	// Set the color
	sprite.Color = rl.White

	// Set the default position and size values
	sprite.Source = physics.NewBody(
		rl.NewVector2(0, 0), rl.NewVector2(float32(texture.Width), float32(texture.Height)),
	)
	sprite.Destination = physics.NewBody(
		rl.NewVector2(0, 0), rl.NewVector2(float32(texture.Width), float32(texture.Height)),
	)

	return sprite
}

func (sprite *Sprite) Render() {
	// If there is no texture, render a colored rectangle
	if sprite.Texture.ID <= 0 {
		RenderRectangle(sprite.Color, sprite.Destination)
	} else {
		// Draw the rectangle with the texture
		RenderTexture(sprite.Texture, sprite.Source, sprite.Destination)
	}
}

func (sprite *Sprite) Destroy() {
	// Unload texture
	DestroyTexture(sprite.Texture)
}
