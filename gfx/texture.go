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

func DestroyTexture(texture rl.Texture2D) {
    // Unload texture for memory
    rl.UnloadTexture(texture)
}

func RenderTexture(texture rl.Texture2D, 
                    srcPosition rl.Vector2, srcSize rl.Vector2,
                    dstPosition rl.Vector2, dstSize rl.Vector2) {
    // Create source rectangle
    srcRectangle := rl.NewRectangle(srcPosition.X, srcPosition.Y, srcSize.X, srcSize.Y)

    // Create destination rectangle
    dstRectangle := rl.NewRectangle(dstPosition.X, dstPosition.Y, dstSize.X, dstSize.Y)

    // Set origin to default constructor (no need to modify it)
    origin := rl.NewVector2(0, 0)

	// Rotation in degrees
	rotation := float32(0)

	// Color
	color := rl.White

    rl.DrawTexturePro(texture, srcRectangle, dstRectangle, origin, rotation, color)
}
