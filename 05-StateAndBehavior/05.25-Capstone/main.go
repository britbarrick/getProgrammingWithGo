// Capstone: Martian Animal Sanctuary
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type turtle struct {
	name string
}

func (t turtle) String() string {
	return t.name
}

func (t turtle) move() string {
	switch rand.Intn(2) {
	case 0:
		return "slowly wanders around"
	default:
		return "lies down slowly"
	}
}

func (t turtle) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "lettuce"
	case 1:
		return "snails"
	default:
		return "minnows"
	}
}

type elephant struct {
	name string
}

func (e elephant) String() string {
	return e.name
}

func (e elephant) move() string {
	switch rand.Intn(2) {
	case 0:
		return "lies down"
	default:
		return "stomps slowly"
	}
}

func (e elephant) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "straw"
	case 1:
		return "peanuts"
	default:
		return "bark"
	}
}

type peacock struct {
	name string
}

func (p peacock) String() string {
	return p.name
}

func (p peacock) move() string {
	switch rand.Intn(2) {
	case 0:
		return "struts about"
	default:
		return "rushes around enclosure"
	}
}

func (p peacock) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "worms"
	case 1:
		return "ants"
	default:
		return "berries"
	}
}

type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v. \n", a, a.move())
	default:
		fmt.Printf("%v eats %v\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	rand.Seed(time.Now().UnixNano())

	animals := []animal{
		turtle{name: "Turtie"},
		elephant{name: "Ellie"},
		peacock{name: "Pete"},
	}

	var sol, hour int

	for {
		fmt.Printf("%2d:00", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("The animals are sleeping.")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}

		time.Sleep(500 * time.Millisecond)

		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}

}
