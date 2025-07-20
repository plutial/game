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
	defer world.Destroy()

	// Set the target frame rate
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		// Updating
		world.Update()

		// Render entities
		world.Render()
	}
}
