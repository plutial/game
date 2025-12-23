package main

import (
	"fmt"

	// Ebitengine
	"github.com/hajimehoshi/ebiten/v2"

	// Game packages
	"github.com/plutial/game/ecs"
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
	"github.com/plutial/game/world"
)

type Game struct {
	Manager ecs.Manager

	// Screen size
	ScreenWidth, ScreenHeight int
}

func NewGame(width, height int, title string) Game {
	game := Game{}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	// Screen size
	game.ScreenWidth = width
	game.ScreenHeight = height

	// Create the game world
	game.Manager = ecs.NewManager()

	// Sprite for renderingame
	ecs.RegisterComponent[gfx.Sprite](&game.Manager)

	// Physics components
	ecs.RegisterComponent[physics.Body](&game.Manager)
	ecs.RegisterComponent[physics.Force](&game.Manager)

	// Entity traits
	ecs.RegisterComponent[physics.Jump](&game.Manager)

	// Tags
	ecs.RegisterComponent[world.PlayerTag](&game.Manager)
	ecs.RegisterComponent[world.EnemyTag](&game.Manager)
	ecs.RegisterComponent[world.TileTag](&game.Manager)
	ecs.RegisterComponent[world.ProjectileTag](&game.Manager)

	// Load maps
	world.LoadMap(&game.Manager, "assets/maps/map0.json")

	// Create the enemies
	world.NewEnemy(&game.Manager)

	// Create the player
	world.NewPlayer(&game.Manager)

	return game
}

func (game *Game) Run() {
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func (game *Game) Update() error {
	// Updating
	game.Manager.Update()

	// Take in input and change it to movement
	world.UpdateMovement(&game.Manager)

	// Attacking
	world.EntityAttack(&game.Manager)

	// Charging
	world.EntityCharge(&game.Manager)

	// Update the physics world
	world.UpdatePhysics(&game.Manager)

	// Update the sprite after all the physics calculations have finished
	world.UpdateSprite(&game.Manager)

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	// Update the screen
	*gfx.GetScreen() = screen

	// Render entities
	game.Manager.Render()
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return game.ScreenWidth, game.ScreenHeight
}
