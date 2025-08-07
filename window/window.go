package window

import (
	// Ebitengine
	"github.com/hajimehoshi/ebiten/v2"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/gfx"
)

// Global game variable
var g Game

type Game struct {
	World ecs.World
}

func Init(width, height int, title string) {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	// Create the game world
	g.World = ecs.NewWorld()
}

func Run() {
	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

func (game *Game) Update() error {
	// Updating
	game.World.Update()

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	// Update the screen
	*gfx.GetScreen() = screen

	// Render entities
	game.World.Render()
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 450
}
