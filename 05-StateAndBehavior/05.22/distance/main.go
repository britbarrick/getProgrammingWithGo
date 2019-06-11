// Go's Got No Class -- Distance
package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

// decimal converts a d/m/s coordinate to dd
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// create new location
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

// convert degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// distance calculation using the Spherical Law of Cosines
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

var (
	earth = world{radius: 6371.0}
	mars  = world{radius: 3389.5}
)

func main() {
	spirit := newLocation(coordinate{14.0, 34.0, 6.2, 'S'}, coordinate{175.0, 28.0, 21.5, 'E'})
	opportunity := newLocation(coordinate{1.0, 56.0, 46.3, 'S'}, coordinate{354.0, 28.0, 24.2, 'E'})
	curiosity := newLocation(coordinate{4.0, 35.0, 22.2, 'S'}, coordinate{137.0, 26.0, 30.1, 'E'})
	inSight := newLocation(coordinate{4.0, 30.0, 0.0, 'N'}, coordinate{135.0, 54.0, 0.0, 'E'})

	fmt.Printf("%.2f, %.2f, %.2f, %.2f, %.2f, %.2f\n",
		mars.distance(spirit, opportunity),
		mars.distance(spirit, curiosity),
		mars.distance(spirit, inSight),
		mars.distance(opportunity, curiosity),
		mars.distance(opportunity, inSight),
		mars.distance(curiosity, inSight))

	london := newLocation(coordinate{51.0, 30.0, 0.0, 'N'}, coordinate{0.0, 8.0, 0.0, 'W'})
	paris := newLocation(coordinate{48.0, 51.0, 0.0, 'N'}, coordinate{2.0, 21.0, 0.0, 'E'})
	fmt.Printf("The distance between London and Paris is %.2f km\n", earth.distance(london, paris))

	stLouis := newLocation(coordinate{38.0, 37.0, 37.2108, 'N'}, coordinate{90.0, 11.0, 57.8472, 'W'})
	washingtonDC := newLocation(coordinate{38.0, 53.0, 23.7516, 'N'}, coordinate{77.0, 0.0, 32.4108, 'W'})
	fmt.Printf("The distance between St. Louis, MO and Washington D.C. is %.2f km\n", earth.distance(stLouis, washingtonDC))

	mtSharp := newLocation(coordinate{5.0, 4.0, 48.0, 'S'}, coordinate{137.0, 51.0, 0.0, 'E'})
	olympusMons := newLocation(coordinate{18.0, 39.0, 0.0, 'N'}, coordinate{226.0, 12.0, 0.0, 'E'})
	fmt.Printf("The distance between Mount Sharp and Olympus Mons is %.2f km\n", mars.distance(mtSharp, olympusMons))
}
