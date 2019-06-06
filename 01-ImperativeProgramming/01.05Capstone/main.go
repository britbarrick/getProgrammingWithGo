// "Capstone: Ticket To Mars"
package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 86400

func main() {
	distance := 62100000
	spaceline := ""
	tripType := ""

	fmt.Println("Spaceline       Days  Trip type  Price")
	fmt.Println("======================================")

	for count := 10; count > 0; count-- {
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}

		speed := rand.Intn(15) + 16
		tripLength := distance / speed / secondsPerDay
		price := speed + 20.0

		if rand.Intn(2) == 1 {
			tripType = "Round-trip"
			price = price * 2
		} else {
			tripType = "One-way"
		}
		fmt.Printf("%-16v %4v %-10v $%4v \n", spaceline, tripLength, tripType, price)
	}
}
