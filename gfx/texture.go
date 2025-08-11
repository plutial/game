package gfx

import (
	"image"
	"image/color"
	"log"

	// Ebitengine
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	// Game packages
	"github.com/plutial/game/physics"
)

// Global screen
var screen *ebiten.Image

// Get the address of the global screen
func GetScreen() **ebiten.Image {
	if screen == nil {
		screen = ebiten.NewImage(800, 450)
	}

	return &screen
}

func NewTexture(path string) *ebiten.Image {
	// Load an image
	texture, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
		panic("Texture was not properly loaded. Image path: " + path)
	}

	return texture
}

func RenderRectangle(color color.RGBA, destinationBody physics.Body, rotation float64) {
	// Options provided by Ebitengine for drawing
	// The order is important for all scales and transformations!
	options := &ebiten.DrawImageOptions{}

	// Apply the position and the size
	// Assuming that the original is a 1x1 rectangle
	options.GeoM.Scale(destinationBody.Size.X, destinationBody.Size.Y)
	options.GeoM.Translate(destinationBody.Position.X, destinationBody.Position.Y)

	// Rotating the image
	// Don't rotate for the angles which are multiples of 360 as rotations are computationally expensive
	if int(rotation)%360 != 0 {
		// Get the center of the image
		center := destinationBody.Center()

		// Translate to the center of the image
		options.GeoM.Translate(-center.X, -center.Y)

		// Rotate around the center
		// The unit for the rotation is in radians
		options.GeoM.Rotate(rotation)

		// Translate back to the screen position
		options.GeoM.Translate(center.X, center.Y)
	}

	// Draw the rectangle with the said color
	coloredTexture := ebiten.NewImage(1, 1)
	coloredTexture.Fill(color)

	screen.DrawImage(coloredTexture, options)
}

func RenderTexture(texture *ebiten.Image,
	sourceBody, destinationBody physics.Body,
	rotation float64,
) {
	// Crop the texture
	sourceRectangle := image.Rect(
		int(sourceBody.Position.X), int(sourceBody.Position.Y),
		int(sourceBody.Position.X)+int(sourceBody.Size.X),
		int(sourceBody.Position.Y)+int(sourceBody.Size.Y),
	)

	subImage := texture.SubImage(sourceRectangle).(*ebiten.Image)

	// Options provided by Ebitengine for drawing
	// The order is important for all scales and transformations!
	options := &ebiten.DrawImageOptions{}

	// Apply the position and the size
	// The division shrinks the image to a 1x1 rectanlge
	options.GeoM.Scale(
		destinationBody.Size.X/sourceBody.Size.X,
		destinationBody.Size.Y/sourceBody.Size.Y,
	)
	options.GeoM.Translate(destinationBody.Position.X, destinationBody.Position.Y)

	// Rotating the image
	// Don't rotate for the angles which are multiples of 360 as rotations are computationally expensive
	if int(rotation)%360 != 0 {
		// Get the center of the image
		center := destinationBody.Center()

		// Translate to the center of the image
		options.GeoM.Translate(-center.X, -center.Y)

		// Rotate around the center
		// The unit for the rotation is in radians
		options.GeoM.Rotate(rotation)

		// Translate back to the screen position
		options.GeoM.Translate(center.X, center.Y)
	}

	// Render the image
	screen.DrawImage(subImage, options)
}
