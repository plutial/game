package main

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/gfx"
)

func main() {
    // Init window
    rl.InitWindow(800, 450, "Game")
	defer rl.CloseWindow()

	sprite := gfx.NewSprite("assets/res/image.png")
	defer sprite.Destroy()

    rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		// Updating
		        
		// Rendering
        rl.BeginDrawing()

        // Clear renderer
		rl.ClearBackground(rl.RayWhite)

		// Render sprite
		sprite.Render()

        // End renderering and swap buffers
		rl.EndDrawing()
	}
}
