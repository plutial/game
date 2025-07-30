package util

import (
	"log"
)

const (
	SPARSE_SET_PAGES = 10
)

type SparseSet[T any] struct {
	sparse        [][]int
	denseToSparse []int
	dense         []T
}

func NewSparseSet[T any]() SparseSet[T] {
	set := SparseSet[T]{}

	set.sparse = make([][]int, 0)
	set.denseToSparse = make([]int, 0)
	set.dense = make([]T, 0)

	return set
}

func (set *SparseSet[T]) Add(index int, value T) {
	if index < 0 {
		log.Fatal("Invalid sparse set index")
	}

	// If index is out of range, double the sparse set
	for index >= len(set.sparse)*SPARSE_SET_PAGES {
		// Create a buffer with the same size as the current sparse set
		buffer := make([][]int, len(set.sparse))

		// If the size of the sparse set is zero, create a buffer with size one
		if len(set.sparse) == 0 {
			buffer = make([][]int, 1)
		}

		// Change all the elements of the buffer to nil
		for i := range buffer {
			buffer[i] = nil
		}

		// Expand the size of the sparse set
		set.sparse = append(set.sparse, buffer...)
	}

	// Calculate the page and the position within the page
	page := index / SPARSE_SET_PAGES
	position := index % SPARSE_SET_PAGES

	if set.sparse[page] == nil {
		set.sparse[page] = make([]int, SPARSE_SET_PAGES)

		for i := range set.sparse[index/SPARSE_SET_PAGES] {
			set.sparse[page][i] = -1
		}
	}

	// Point the sparse set's index to the dense set's index
	set.sparse[page][position] = len(set.dense)

	// Create a pointer to the dense set to the sparse set
	set.denseToSparse = append(set.denseToSparse, index)

	// Add the value to the dense set
	set.dense = append(set.dense, value)
}

func (set *SparseSet[T]) Set(index int, value T) {
	// Calculate the page and the position within the page
	page := index / SPARSE_SET_PAGES
	position := index % SPARSE_SET_PAGES

	// Point the sparse set's index to the dense set's index
	denseIndex := set.sparse[page][position]

	// Add the value to the dense set
	set.dense[denseIndex] = value
}

func (set *SparseSet[T]) GetAddress(index int) (*T, bool) {
	// Calculate the page and the position within the page
	page := index / SPARSE_SET_PAGES
	position := index % SPARSE_SET_PAGES

	// Check to see if the index is out of bounds
	if len(set.sparse)*SPARSE_SET_PAGES <= index {
		return nil, false
	}

	// Point the sparse set's index to the dense set's index
	if set.sparse[page] == nil {
		return nil, false
	}

	if set.sparse[page][position] == -1 {
		return nil, false
	}

	denseIndex := set.sparse[page][position]

	// If the dense index somehow is larger than the length of the slice
	if denseIndex >= len(set.dense) {
		panic("The dense index is greater than or equal to the length of the slice.")
	}

	// Add the value to the dense set
	return &set.dense[denseIndex], true
}

func (set *SparseSet[T]) Get(index int) (T, bool) {
	valueAddress, ok := set.GetAddress(index)

	// If the value addres is nil ie. something went wrong
	if valueAddress == nil {
		var temp T
		return temp, ok
	}

	return *valueAddress, ok
}

func (set *SparseSet[T]) Delete(index int) {
	// Check if the index is valid in the first place
	_, ok := set.Get(index)
	if !ok {
		return
	}

	// Check to see if the index is out of bounds
	if len(set.sparse)*SPARSE_SET_PAGES <= index {
		panic("Index out of bounds to remove.")
	}

	// Calculate the page and the position within the page
	page := index / SPARSE_SET_PAGES
	position := index % SPARSE_SET_PAGES

	// Point the sparse set's index to the dense set's index
	denseIndex := set.sparse[page][position]

	// Set the last element of the dense set to the dense index
	lastIndex := len(set.dense) - 1
	set.dense[denseIndex] = set.dense[lastIndex]

	// Change the dense to sparse pointer
	set.denseToSparse[denseIndex] = set.denseToSparse[lastIndex]

	// Change the sparse to dense pointer
	swappedIndex := set.denseToSparse[denseIndex]

	swappedPage := swappedIndex / SPARSE_SET_PAGES
	swwappedPosition := swappedIndex % SPARSE_SET_PAGES

	set.sparse[swappedPage][swwappedPosition] = index

	// Change the original sparse to dense pointer to nil
	// This should happen AFTER the swapped sparse pointer is corrected
	// Otherwise, if the index is pointing at the last element,
	// the sparse set isn't properly changed
	set.sparse[page][position] = -1

	// Remove the last elements of the dense sets
	set.dense = set.dense[:lastIndex]
	set.denseToSparse = set.denseToSparse[:lastIndex]
}
