package gfx

import (
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

func RenderRectanlge(color rl.Color,
	dstPosition, dstSize rl.Vector2,
) {
	// Create a sample rectangle to make the code easier to read
	dstRectangle := rl.NewRectangle(dstPosition.X, dstPosition.Y, dstSize.X, dstSize.Y)

	// Scale the destination rectangle to fit window size
	scale := rl.NewVector2(float32(rl.GetScreenWidth())/800, float32(rl.GetScreenHeight())/450)

	dstRectangle.X *= scale.X
	dstRectangle.Width *= scale.X

	dstRectangle.Y *= scale.Y
	dstRectangle.Height *= scale.Y

	// Draw the rectangle with the said color
	rl.DrawRectangleRec(dstRectangle, color)
}

func RenderTexture(texture rl.Texture2D,
	srcPosition, srcSize rl.Vector2,
	dstPosition, dstSize rl.Vector2,
) {
	// Create source rectangle
	srcRectangle := rl.NewRectangle(srcPosition.X, srcPosition.Y, srcSize.X, srcSize.Y)

	// Create destination rectangle
	dstRectangle := rl.NewRectangle(dstPosition.X, dstPosition.Y, dstSize.X, dstSize.Y)

	// Scale the destination rectangle to window size
	scale := rl.NewVector2(float32(rl.GetScreenWidth())/800, float32(rl.GetScreenHeight())/450)

	dstRectangle.X *= scale.X
	dstRectangle.Width *= scale.X

	dstRectangle.Y *= scale.Y
	dstRectangle.Height *= scale.Y

	// Set origin to default constructor (no need to modify it)
	origin := rl.NewVector2(0, 0)

	// Rotation in degrees
	rotation := float32(0)

	// Color
	color := rl.White

	rl.DrawTexturePro(texture, srcRectangle, dstRectangle, origin, rotation, color)
}

func DestroyTexture(texture rl.Texture2D) {
	// Unload texture for memory
	rl.UnloadTexture(texture)
}
