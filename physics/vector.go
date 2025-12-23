package physics

import (
	"fmt"
	"math"
)

type Vector2f struct {
	X, Y float64
}

func NewVector2f(x, y float64) Vector2f {
	return Vector2f{x, y}
}

// Pretty formatting
func (vector Vector2f) String() string {
	return fmt.Sprintf("%f\t%f", vector.X, vector.Y)
}

// Utilizes pythagoras to find the distance between two points
func (vectorA Vector2f) Distance(vectorB Vector2f) float64 {
	return math.Hypot(vectorA.X+vectorB.X, vectorA.Y+vectorB.Y)
}

// Magnitude of the vector
func (vector Vector2f) Magnitude() float64 {
	return math.Hypot(vector.X, vector.Y)
}

// Slope of the vector
// Used to calculate the gradient of the line
func (vector Vector2f) Slope() float64 {
	return vector.Y / vector.X
}
