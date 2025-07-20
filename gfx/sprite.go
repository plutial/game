package gfx

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	// The texture itself
	Texture rl.Texture2D

	// The color the sprite will render with if the texture is nil
	Color rl.Color

	// The source co-ordinates and size
	SrcPosition, SrcSize rl.Vector2

	// The destination co-ordinates and size
	DstPosition, DstSize rl.Vector2
}

func NewSprite(texture rl.Texture2D) Sprite {
	var sprite Sprite

	// Load the texture
	sprite.Texture = texture

	// Set the color
	sprite.Color = rl.White

	// Set the default position and size values
	sprite.SrcPosition = rl.NewVector2(0, 0)
	sprite.SrcSize = rl.NewVector2(16.0, 16.0)

	sprite.DstPosition = rl.NewVector2(0, 0)
	sprite.DstSize = rl.NewVector2(16.0, 16.0)

	return sprite
}

func (sprite *Sprite) Render() {
	// If there is no texture, render a colored rectangle
	if sprite.Texture.ID <= 0 {
		RenderRectanlge(sprite.Color, sprite.DstPosition, sprite.DstSize)
	} else {
		// Draw the rectangle with the texture
		RenderTexture(sprite.Texture, sprite.SrcPosition, sprite.SrcSize, sprite.DstPosition, sprite.DstSize)
	}
}

func (sprite *Sprite) Destroy() {
	// Unload texture
	DestroyTexture(sprite.Texture)
}
