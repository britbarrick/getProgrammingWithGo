// "Whole Numbers"
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	balance := 0
	for balance < 2000 {
		cents := balance % 100
		dollars := (balance - cents) / 100
		switch rand.Intn(3) {
		case 0:
			balance += 5
		case 1:
			balance += 10
		case 2:
			balance += 25
		}
		fmt.Printf("$ %d.%.2d\n", dollars, cents)
	}

}
