package main

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// "fmt"
	// "reflect"
	// "log"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

func main() {
	// Init window
	rl.InitWindow(800, 450, "Game")
	defer rl.CloseWindow()

	// Create the game world
	world := ecs.NewWorld()

	// Register components
	ecs.RegisterComponent[gfx.Sprite](&world)
	ecs.RegisterComponent[physics.Body](&world)
	ecs.RegisterComponent[physics.Force](&world)

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
}
