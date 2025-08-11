package ecs

// Does the entity exist
type Alive bool

// Check if an entity is alive
func (manager *Manager) IsEntityAlive(id int) bool {
	componentSet := *GetComponentSet[Alive](manager)

	_, ok := componentSet.Get(id)

	return ok
}

// Create an entity
func (manager *Manager) NewEntity() int {
	for id := range manager.Size {
		// If the entity is not alive, assign the new entity id
		if !manager.IsEntityAlive(id) {
			// Check the entity is now alive
			componentSet := GetComponentSet[Alive](manager)
			componentSet.Add(id, true)

			return id
		}
	}

	// If every entity that currently exists is alive, add a new entity position
	id := manager.Size

	// Check the entity is now alive
	componentSet := GetComponentSet[Alive](manager)
	componentSet.Add(id, true)

	// Increase the number of entities
	manager.Size++

	return id
}

// Delete an entity
func (manager *Manager) DeleteEntity(id int) {
	// Check that the entity is alive before adding it to the delete table
	if !manager.IsEntityAlive(id) {
		return
	}

	// Add the entity to the list to remove
	manager.ToDelete = append(manager.ToDelete, id)
}
