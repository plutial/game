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

func RenderRectangle(color color.RGBA, destionationBody physics.Body) {
	// Options provided by Ebitengine for drawing
	options := &ebiten.DrawImageOptions{}

	// Apply the position and the size
	// Assuming that the original is a 1x1 rectangle
	// The order is important!
	options.GeoM.Scale(destionationBody.Size.X, destionationBody.Size.Y)
	options.GeoM.Translate(destionationBody.Position.X, destionationBody.Position.Y)

	// Scale the destination rectangle to fit window size
	/*scale := physics.NewVector2(float64(rl.GetScreenWidth())/800, float64(rl.GetScreenHeight())/450)

	destination.X *= scale.X
	destination.Width *= scale.X

	destination.Y *= scale.Y
	destination.Height *= scale.Y*/

	// Draw the rectangle with the said color
	coloredTexture := ebiten.NewImage(1, 1)
	coloredTexture.Fill(color)

	screen.DrawImage(coloredTexture, options)
}

func RenderTexture(texture *ebiten.Image,
	sourceBody, destinationBody physics.Body,
) {
	// Scale the destination rectangle to window size
	/*scale := physics.NewVector2(float32(rl.GetScreenWidth())/800, float32(rl.GetScreenHeight())/450)

	destination.X *= scale.X
	destination.Width *= scale.X

	destination.Y *= scale.Y
	destination.Height *= scale.Y*/

	// Crop the texture
	sourceRectangle := image.Rect(
		int(sourceBody.Position.X), int(sourceBody.Position.Y),
		int(sourceBody.Position.X)+int(sourceBody.Size.X),
		int(sourceBody.Position.Y)+int(sourceBody.Size.Y),
	)

	subImage := texture.SubImage(sourceRectangle).(*ebiten.Image)

	// Apply the destination body
	// Assuming that the dimensions of the image are the same as the cropped image
	// The order is important!
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(
		destinationBody.Size.X/sourceBody.Size.X,
		destinationBody.Size.Y/sourceBody.Size.Y,
	)
	options.GeoM.Translate(
		destinationBody.Position.X, destinationBody.Position.Y,
	)

	// Render the image
	screen.DrawImage(subImage, options)
}
