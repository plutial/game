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

	// The rotation of the sprite in radians
	Rotation float64

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

	// Set the rotation in radians
	sprite.Rotation = 0

	// Set the default position and size values
	// The sizes will be the size of the image
	sprite.Source = physics.NewBody(
		physics.NewVector2(0, 0),
		physics.NewVector2(
			float64(sprite.Image.Bounds().Dx()),
			float64(sprite.Image.Bounds().Dy()),
		),
	)
	sprite.Destination = physics.NewBody(
		physics.NewVector2(0, 0),
		physics.NewVector2(
			float64(sprite.Image.Bounds().Dx()),
			float64(sprite.Image.Bounds().Dy()),
		),
	)

	return sprite
}

func (sprite *Sprite) Render() {
	// If there is no texture, render a colored rectangle
	if sprite.Image == nil {
		RenderRectangle(sprite.Color, sprite.Destination, sprite.Rotation)
	} else {
		// Draw the rectangle with the texture
		RenderTexture(sprite.Image, sprite.Source, sprite.Destination, sprite.Rotation)
	}
}

func (sprite *Sprite) Destroy() {
	// TODO: Unload texture
}
