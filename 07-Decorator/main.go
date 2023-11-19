/*
	This pattern is used to add additional functionality to the existing types.
	It obey the Open Close principle.
*/

package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius int
}

func (c *Circle) Render() string {
	return "Rendering Circle"
}

func (c *Circle) Resize() string {
	return "Resizing Circle"
}

type Square struct {
	Side int
}

func (s *Square) Render() string {
	return "Rendering Square"
}

// Now we want to add a property named color but we don't want to modify the existing type.
// Here we will use decorator pattern and create a new struct ColoredShape

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	// Using underlying Render method from Shape interface
	return c.Shape.Render() + " " + c.Color
}

func main() {
	circle := &Circle{}
	fmt.Println(circle.Render())

	coloredCircle := &ColoredShape{
		Shape: circle,
		Color: "Blue",
	}

	fmt.Println(coloredCircle.Render())
}
