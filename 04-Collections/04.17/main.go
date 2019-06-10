// Slices: Windows into Arrays
package main

import "fmt"

type planets []string

func (p planets) terraform() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func main() {
	var planets planets = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptue"}

	planets[3:4].terraform()
	planets[6:].terraform()
	fmt.Println(planets)
}
