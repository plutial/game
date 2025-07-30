package gfx

import (
	"image/color"

	// Ebitengine
	"github.com/hajimehoshi/ebiten/v2"

	// Game packages
	"github.com/plutial/game/physics"
)

type Sprite struct {
	// The image itself
	Image *ebiten.Image

	// The color the sprite will render with if the texture is nil
	Color color.RGBA

	// The source co-ordinates and size
	Source physics.Body

	// The destination co-ordinates and size
	Destination physics.Body
}

func NewSprite(texture *ebiten.Image) Sprite {
	var sprite Sprite

	// The image (renamed to texture to avoid clash with the image package)
	sprite.Image = texture

	// Set the color (white)
	sprite.Color = color.RGBA{255, 255, 255, 255}

	// Set the default position and size values
	sprite.Source = physics.NewBody(
		physics.NewVector2(0, 0), physics.NewVector2(16, 16),
	)
	sprite.Destination = physics.NewBody(
		physics.NewVector2(0, 0), physics.NewVector2(16, 16),
	)

	return sprite
}

func (sprite *Sprite) Render() {
	// If there is no texture, render a colored rectangle
	if sprite.Image == nil {
		RenderRectangle(sprite.Color, sprite.Destination)
	} else {
		// Draw the rectangle with the texture
		RenderTexture(sprite.Image, sprite.Source, sprite.Destination)
	}
}

func (sprite *Sprite) Destroy() {
	// TODO: Unload texture
}
