package ecs

import (
	"log"
	"fmt"
	"reflect"
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
	// Gravity
	Gravity 	float32
	MaxGravity 	float32

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

	// Gravity
	world.Gravity = 0.1
	world.MaxGravity = 5 

	// ECS
	world.EntityHasComponent = make(map[int]map[reflect.Type]bool)
	world.ComponentPool = make(map[reflect.Type]any)
		
	// Register components
	RegisterComponent[gfx.Sprite](&world)
	RegisterComponent[physics.Body](&world)
	RegisterComponent[physics.Force](&world)
	RegisterComponent[physics.Collisions](&world)

	// Player traits
	RegisterComponent[Jump](&world)

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

	collisions, err := GetComponent[physics.Collisions](world, playerId)

	if err != nil {
		log.Fatal(err)
	}

	jump, err := GetComponent[Jump](world, playerId)

	if err != nil {
		log.Fatal(err)
	}

	// Horizontal movement
	// speed := float32(3)

	if rl.IsKeyDown(rl.KeyA) {
		force.Acceleration.X = max(-3, force.Acceleration.X - 0.3)
	}

	if rl.IsKeyDown(rl.KeyD) {
		force.Acceleration.X = min(3, force.Acceleration.X + 0.3)
	} 

	if !rl.IsKeyDown(rl.KeyA) && !rl.IsKeyDown(rl.KeyD) {
		if force.Acceleration.X < 0 {
			force.Acceleration.X = min(0, force.Acceleration.X + 0.6)
		} else {
			force.Acceleration.X = max(0, force.Acceleration.X - 0.6)
		}
	}


	// Jump
	if collisions.Down {
		// How long the body has been in the air for
		jump.AirTime = 0
		
		// Numbers of jumps possible at a time
		// 2 for double jump, 3 for triple jump ...
		// Currently not working but I'm not bother to fix it
		jump.Jumps = 1
	} else {
		// If the body is not touching the ground, it's in the air
		jump.AirTime += 1
	}

	if rl.IsKeyPressed(rl.KeyW) {
		// Register a jump 
		// A jump can be registered even if the body has not yet touched the ground
		jump.JumpRegistered = 3

		// Fix to enable multiple jumps
		// Currently not working, coyote time doesn't work after using this fix
		// jump.AirTime = 0
	}

		
	// If the body does not hit the ground in time, it won't jump
	if jump.JumpRegistered > 0 {
		// If the body can jump, and if it is on the ground (kind of... coyote time)
		if jump.Jumps > 0 && jump.AirTime < 5 {
			// How high it goes (and the actual jump part)
			force.Acceleration.Y = -3

			// Take off an available jump
			jump.Jumps -= 1

			// Stop registering the jumps
			jump.JumpRegistered = 0
		} else {
			// Tick down the timer of the register
			jump.JumpRegistered -= 1
		}
	}
}

// Update tile physics against a body with force
func (world *World) UpdateTilePhysics(body *physics.Body, force *physics.Force, collisions *physics.Collisions, tiles []int) {
	// Slice to store tiles which have collided
	type TileCollisionData struct {
		TileId 		int
		Distance 	float32
	}

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

	// Reset the collisions (if the entity has the collisions component)
	if collisions != nil {
		*collisions = physics.Collisions {}
	}

	// Resolve the collisions
	for _, data := range tileCollisionData {
		tileId := data.TileId

		tileBody, err := GetComponent[physics.Body](world, tileId)

		if err != nil {
			log.Fatal(err)
		}

		collision, velocityResolve, contactNormal := physics.BodyDynamicVsBodyResolve(*body, *tileBody, force.Velocity)

		if collision {
			// Update the collision velocity
			force.Velocity = velocityResolve

			// Update the collision direction
			if collisions != nil {
				collisions.Update(contactNormal)
			}
		}
	}
}

// Update all the entites with a body and force
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

		var collisions *physics.Collisions
		if HasComponent[physics.Collisions](world, id) {
			collisions, err = GetComponent[physics.Collisions](world, id)

			if err != nil {
				log.Fatal(err)
			}
		}

		// Apply gravity
		force.Acceleration.Y = min(world.MaxGravity, force.Acceleration.Y + world.Gravity)

		// Update acceleration
		force.Velocity.X += force.Acceleration.X
		force.Velocity.Y += force.Acceleration.Y

		// Handle tile collisions
		// This MUST be handled at the end AFTER acceleration has been applied
		world.UpdateTilePhysics(body, force, collisions, tiles)

		// Update the body position
		body.Position.X += force.Velocity.X
		body.Position.Y += force.Velocity.Y

		fmt.Println(force.Acceleration)
		fmt.Println(force.Velocity)
		fmt.Println("")

		// Reset the velocity after calculation
		force.Velocity.X = 0
		force.Velocity.Y = 0
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
