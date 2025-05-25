package main

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/ecs"
)

func main() {
	// Disable logging
	rl.SetTraceLogLevel(rl.LogWarning)

	// Init window
	rl.InitWindow(800, 450, "Game")
	defer rl.CloseWindow()

	// Create the game world
	world := ecs.NewWorld()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		// Updating
		world.Update()

		// Rendering
		rl.BeginDrawing()

		// Clear renderer with a white background
		rl.ClearBackground(rl.RayWhite)

		// Render entities
		world.Render()

		// End renderering and swap buffers
		rl.EndDrawing()
	}

	world.Destroy()
}
