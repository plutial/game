package main

import (
	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
	//"fmt"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/gfx"
	/*"github.com/plutial/game/physics"*/
)

func main() {
	// Init window
	rl.InitWindow(800, 450, "Game")
	defer rl.CloseWindow()

	world := ecs.NewWorld()
	ecs.RegisterComponent[gfx.Sprite](&world)

	//texture := gfx.NewTexture("assets/res/image.png")
	//sprite := gfx.NewSprite(texture)
	entity := uint32(0)
	ecs.AddComponent[gfx.Sprite](&world, entity)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		// Updating

		// Rendering
		rl.BeginDrawing()

		// Clear renderer as white background
		rl.ClearBackground(rl.RayWhite)

		// Render entities

		// End renderering and swap buffers
		rl.EndDrawing()
	}
}
