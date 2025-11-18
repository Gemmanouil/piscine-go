package main

import (
	"fmt"
)

// Define a struct named 'point' with two integer fields: x and y
type point struct {
	x int
	y int
}

// This function sets the values of x and y in the given point to 42 and 21
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	// Create a pointer to a new point struct
	points := &point{}

	// Call the function to set the values of the point
	setPoint(points)

	// Print the values of x and y using formatted output
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
