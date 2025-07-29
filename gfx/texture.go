package gfx

import (
	// Game packages
	"github.com/plutial/game/physics"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewTexture(path string) rl.Texture2D {
	// Load texture
	texture := rl.LoadTexture(path)

	// Check if the texture was properly loaded
	if texture.ID <= 0 {
		panic("Texture was not properly loaded. Image path: " + path)
	}

	// Return texture
	return texture
}

func RenderRectangle(color rl.Color, destionationBody physics.Body) {
	// Create a sample rectangle to make the code easier to read
	destination := destionationBody.Rectangle()

	// Scale the destination rectangle to fit window size
	scale := rl.NewVector2(float32(rl.GetScreenWidth())/800, float32(rl.GetScreenHeight())/450)

	destination.X *= scale.X
	destination.Width *= scale.X

	destination.Y *= scale.Y
	destination.Height *= scale.Y

	// Draw the rectangle with the said color
	rl.DrawRectangleRec(destination, color)
}

func RenderTexture(texture rl.Texture2D,
	sourceBody, destinationBody physics.Body,
) {
	// Create source rectangle
	source := sourceBody.Rectangle()

	// Create destination rectangle
	destination := destinationBody.Rectangle()

	// Scale the destination rectangle to window size
	scale := rl.NewVector2(float32(rl.GetScreenWidth())/800, float32(rl.GetScreenHeight())/450)

	destination.X *= scale.X
	destination.Width *= scale.X

	destination.Y *= scale.Y
	destination.Height *= scale.Y

	// Set origin to default constructor (no need to modify it)
	origin := rl.NewVector2(0, 0)

	// Rotation in degrees
	rotation := float32(0)

	// Color
	color := rl.White

	rl.DrawTexturePro(texture,
		source, destination,
		origin, rotation, color,
	)
}

func DestroyTexture(texture rl.Texture2D) {
	// Unload texture for memory
	rl.UnloadTexture(texture)
}
