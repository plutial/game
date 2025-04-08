package gfx

import (
	// Raylib
    rl "github.com/gen2brain/raylib-go/raylib"
)

func TextureInit(path string) rl.Texture2D {
    // Load texture
    texture := rl.LoadTexture(path)

    // Check if the texture was properly loaded
    if texture.ID <= 0 {
        panic("Texture was not properly loaded. Image path: " + path)
    }

    // Return texture
    return texture
}

func TextureDestroy(texture rl.Texture2D) {
    // Unload texture for memory
    rl.UnloadTexture(texture)
}

func TextureRender(texture rl.Texture2D, 
                    src_pos rl.Vector2, src_size rl.Vector2,
                    dst_pos rl.Vector2, dst_size rl.Vector2) {
    // Create source rectangle
    src_rect := rl.NewRectangle(src_pos.X, src_pos.Y, src_size.X, src_size.Y)

    // Create destination rectangle
    dst_rect := rl.NewRectangle(dst_pos.X, dst_pos.Y, dst_size.X, dst_size.Y)

    // Set origin to default constructor (no need to modify it)
    origin := rl.NewVector2(0, 0)

    rl.DrawTexturePro(texture, src_rect, dst_rect, origin, 0.0, rl.White)
}
