package gfx

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	// The texture itself
	Texture rl.Texture2D

	// The source co-ordinates and size
	SrcPosition rl.Vector2
	SrcSize     rl.Vector2

	// The destination co-ordinates and size
	DstPosition rl.Vector2
	DstSize     rl.Vector2
}

func NewSprite(path string) Sprite {
    var sprite Sprite

	// Load texture
	sprite.Texture = TextureInit(path)

	// Set default position and size values
	sprite.SrcPosition = rl.NewVector2(0, 0)
	sprite.SrcSize = rl.NewVector2(16.0, 16.0)

	sprite.DstPosition = rl.NewVector2(0, 0)
	sprite.DstSize = rl.NewVector2(16.0 * 4, 16.0 * 4)

    return sprite
}

func (sprite *Sprite) Init(path string) {
    *sprite = NewSprite(path)
}

func (sprite *Sprite) Destroy() {
	// Unload texture
	TextureDestroy(sprite.Texture)
}

func (sprite *Sprite) Render() {
	TextureRender(sprite.Texture, sprite.SrcPosition, sprite.SrcSize, sprite.DstPosition, sprite.DstSize)
}
