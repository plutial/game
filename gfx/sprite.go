package gfx

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	// The texture itself
	Texture rl.Texture2D

	// The source co-ordinates and size
	SrcPosition, SrcSize rl.Vector2

	// The destination co-ordinates and size
	DstPosition, DstSize rl.Vector2
}

func NewSprite(texture rl.Texture2D) Sprite {
    var sprite Sprite

	// Load texture
	sprite.Texture = texture

	// Set default position and size values
	sprite.SrcPosition = rl.NewVector2(0, 0)
	sprite.SrcSize = rl.NewVector2(16.0, 16.0)

	sprite.DstPosition = rl.NewVector2(0, 0)
	sprite.DstSize = rl.NewVector2(16.0 * 4, 16.0 * 4)

    return sprite
}

func (sprite *Sprite) Destroy() {
	// Unload texture
	DestroyTexture(sprite.Texture)
}

func (sprite *Sprite) Render() {
	RenderTexture(sprite.Texture, sprite.SrcPosition, sprite.SrcSize, sprite.DstPosition, sprite.DstSize)
}
