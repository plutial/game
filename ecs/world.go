package ecs

import (
	"log"
	"reflect"
	// "math"
	"sort"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"

	// Game packages
	"github.com/plutial/game/gfx"
	"github.com/plutial/game/physics"
)

// Get the type of a generic type interface as reflect.Type variable
func getType[T any]() reflect.Type {
	var temp T
	return reflect.TypeOf(temp)
}

// Contains entity count information,
// which entity contains which components,
// and component storage, stored as slices
type World struct {
	// Entity count
	Size int

	// Check if an entity has a component
	EntityHasComponent map[int]map[reflect.Type]bool

	// Component storage
	ComponentPool map[reflect.Type]any
}

// Create a new world and its entities' components
func NewWorld() World {
	world := World {}
	world.EntityHasComponent = make(map[int]map[reflect.Type]bool)
	world.ComponentPool = make(map[reflect.Type]any)
		
	// Register components
	RegisterComponent[gfx.Sprite](&world)
	RegisterComponent[physics.Body](&world)
	RegisterComponent[physics.Force](&world)

	// Tags
	RegisterComponent[PlayerTag](&world)
	RegisterComponent[TileTag](&world)

	// Load maps
	world.LoadMap("assets/maps/map0.json")

	// Create the player
	world.NewPlayer()

	return world
}

func (world *World) NewEntity() int {
	for id := range world.Size {
		_, ok := world.EntityHasComponent[id]
		
		// If the entity is not alive, assign the new entity id
		if !ok {
			// Check the entity is now alive
			world.EntityHasComponent[id] = make(map[reflect.Type]bool)
			return id 
		}
	}

	// If every entity that currently exists is alive, add a new entity position
	id := world.Size
			
	// Check the entity is now alive
	world.EntityHasComponent[id] = make(map[reflect.Type]bool)

	// Increase the number of entities
	world.Size++ 

	return id
}

func (world *World) DeleteEntity(id int) {
	delete(world.EntityHasComponent, id)
}

func (world *World) UpdateInput() {
	// Get the player id
	entities := GetEntities[PlayerTag](world)
	playerId := entities[0]

	force, err := GetComponent[physics.Force](world, playerId)

	if err != nil {
		log.Fatal(err)
	}

	speed := float32(3)

	if rl.IsKeyDown(rl.KeyA) {
		force.Velocity.X -= speed
	}

	if rl.IsKeyDown(rl.KeyD) {
		force.Velocity.X += speed
	}

	if rl.IsKeyDown(rl.KeyW) {
		force.Velocity.Y -= speed
	}

	if rl.IsKeyDown(rl.KeyS) {
		force.Velocity.Y += speed
	}
}

type TileCollisionData struct {
	TileId 		int
	Distance 	float32
}

func (world *World) UpdatePhysics() {
	// Get all the entities which have the body component and the force component
	entities := GetEntities2[physics.Body, physics.Force](world)

	// Get all tiles
	tiles := GetEntities[TileTag](world)

	for _, id := range entities {
		// Get the components
		body, err := GetComponent[physics.Body](world, id)

		if err != nil {
			log.Fatal(err)
		}

		force, err := GetComponent[physics.Force](world, id)

		if err != nil {
			log.Fatal(err)
		}

		// Handle tile collisions
		// Slice to store tiles which have collided
		tileCollisionData := make([]TileCollisionData, 0)

		// Check which tiles could collided with the body
		for _, tileId := range tiles {
			tileBody, err := GetComponent[physics.Body](world, tileId)

			if err != nil {
				log.Fatal(err)
			}
		
			// Carry out a broad phase to stop handling 
			// Minimize expensive physics on absurd tiles that will never collide with
			collision := physics.BodyBroadPhase(*body, *tileBody, force.Velocity)
			
			if !collision {
				continue
			}

			// Check for collision
			collision, _, _ = physics.BodyDynamicVsBody(*body, *tileBody, force.Velocity)

			if collision {
				// Get the distance from the body to the tile
				distance := physics.GetDistance(body.Position, tileBody.Position)

				data := TileCollisionData { tileId , distance }

				// Add to the collided tile list
				tileCollisionData = append(tileCollisionData, data)
			}
		}

		// Sort the tiles by which tiles are the closest to the body
		// This is a fix to imitate actual physics (and to handle other cases)
		sort.SliceStable(tileCollisionData, func(a, b int) bool {
			return tileCollisionData[a].Distance < tileCollisionData[b].Distance
		})

		// Resolve the collisions
		for _, data := range tileCollisionData {
			tileId := data.TileId

			tileBody, err := GetComponent[physics.Body](world, tileId)

			if err != nil {
				log.Fatal(err)
			}

			collision, velocityResolve := physics.BodyDynamicVsBodyResolve(*body, *tileBody, force.Velocity)

			if collision {
				force.Velocity = velocityResolve
			}
		}

		// Update the body position
		body.Position.X += force.Velocity.X
		body.Position.Y += force.Velocity.Y

		// Reset the velocity after calculation
		force.Velocity = rl.NewVector2(0, 0)
	}
}

func (world *World) UpdateSprite() {
	// Get all the entities which have the sprite component and the body component
	entities := GetEntities2[gfx.Sprite, physics.Body](world)

	for _, id := range entities {
		// Get the components
		sprite, err := GetComponent[gfx.Sprite](world, id)

		if err != nil {
			log.Fatal(err)
		}

		body, err := GetComponent[physics.Body](world, id)
	
		if err != nil {
			log.Fatal(err)
		}

		// Update the position of the sprite
		sprite.DstPosition = body.Position
	}
}

func (world *World) Update() {
	// Take in input
	world.UpdateInput()

	// Update the physics world
	world.UpdatePhysics()

	// Update the sprite after all the physics is finished
	world.UpdateSprite()
}

func (world *World) Render() {
	entities := GetEntities[gfx.Sprite](world)
	
	for _, id := range entities {
		sprite, err := GetComponent[gfx.Sprite](world, id)

		if err != nil {
			log.Fatal(err)
		}

		sprite.Render()
	}
}

func (world *World) Destroy() {
	entities := GetEntities[gfx.Sprite](world)
	
	for _, id := range entities {
		sprite, err := GetComponent[gfx.Sprite](world, id)

		if err != nil {
			log.Fatal(err)
		}

		sprite.Destroy()
	}
}
