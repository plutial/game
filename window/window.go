package window

import (
	// Ebitengine
	"github.com/hajimehoshi/ebiten/v2"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/gfx"
)

type Game struct{}

// Global world
var world ecs.World

func Init(width, height int, title string) {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	// Create the game world
	world = ecs.NewWorld()
}

func Run() {
	game := Game{}
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}

func (game *Game) Update() error {
	// Updating
	world.Update()

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	// Update the screen
	*gfx.GetScreen() = screen

	// Render entities
	world.Render()
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 450
}
