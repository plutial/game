package ecs

// Does the entity exist
type Alive bool

// Check if an entity is alive
func (world *World) IsEntityAlive(id int) bool {
	componentSet := *GetComponentSet[Alive](world)

	_, ok := componentSet.Get(id)

	return ok
}

// Create an entity
func (world *World) NewEntity() int {
	for id := range world.Size {
		// If the entity is not alive, assign the new entity id
		if !world.IsEntityAlive(id) {
			// Check the entity is now alive
			componentSet := GetComponentSet[Alive](world)
			componentSet.Add(id, true)

			return id
		}
	}

	// If every entity that currently exists is alive, add a new entity position
	id := world.Size

	// Check the entity is now alive
	componentSet := GetComponentSet[Alive](world)
	componentSet.Add(id, true)

	// Increase the number of entities
	world.Size++

	return id
}

// Delete an entity
func (world *World) DeleteEntity(id int) {
	// Check that the entity is alive before adding it to the delete table
	if !world.IsEntityAlive(id) {
		return
	}

	// Add the entity to the list to remove
	world.ToDelete = append(world.ToDelete, id)
}
