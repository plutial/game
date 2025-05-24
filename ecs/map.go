package ecs

import (
	"encoding/json"
	"io/ioutil"
	"log"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

type GameMapData struct {
	// Tile size
	TileWidth 	int `json:"tilewidth"`
	TileHeight 	int	`json:"tileheight"`
	
	// Tile Layer Size
	LayerWidth 	int `json:"width"`
	LayerHeight	int `json:"height"`

	// Tile layers
	TileLayers 	[]TileLayerData `json:"layers"`
}

type TileLayerData struct {
	Data []int 	`json:"data"`
	Name string	`json:"name"`
}

type TileTag bool

func (world *World) LoadMap(path string) {
	// Open the json file
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	// Map
	var gameMapData GameMapData
 
	json.Unmarshal(data, &gameMapData)

	// Load the tile texture
	tileTexture := gfx.NewTexture("assets/res/GrassTiles.png")

	// Put the tiles into the world
	for y := range gameMapData.LayerHeight {
		for x := range gameMapData.LayerWidth {
			tileSourceId := gameMapData.TileLayers[0].Data[y * gameMapData.LayerWidth + x]
		
			// If the tile does not exist, then do not load it
			if tileSourceId == 0 {
				continue
			}

			// Create an entity
			id := world.NewEntity()

			// Add the tile tag
			AddComponent[TileTag](world, id)

			// Create the sprite component
			sprite, err := AddComponent[gfx.Sprite](world, id)

			if err != nil {
				log.Fatal(err)
			}

			*sprite = gfx.NewSprite(tileTexture)

			// Source rectangle
			sprite.SrcPosition.X = float32((tileSourceId - 1) % int(tileTexture.Width / 16)) * 16
			sprite.SrcPosition.Y = float32((tileSourceId - 1) / int(tileTexture.Width / 16)) * 16

			sprite.SrcSize = rl.NewVector2(16, 16)

			// Destination rectangle
			sprite.DstPosition = rl.NewVector2(float32(x) * 16, float32(y) * 16)
			sprite.DstSize = rl.NewVector2(16, 16)

			// Create the physics body
			body, err := AddComponent[physics.Body](world, id)

			if err != nil {
				log.Fatal(err)
			}

			position := rl.NewVector2(float32(x) * 16, float32(y) * 16)
			size := rl.NewVector2(16, 16)

			*body = physics.NewBody(position, size)
		}
	}
}
