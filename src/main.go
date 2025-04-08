package main

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
    // Init window
    rl.InitWindow(800, 450, "Game")
	defer rl.CloseWindow()
   
    rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		// Updating
		        
		// Rendering
        rl.BeginDrawing()

        // Clear renderer
		rl.ClearBackground(rl.RayWhite)

        // End renderering and swap buffers
		rl.EndDrawing()
	}
}
