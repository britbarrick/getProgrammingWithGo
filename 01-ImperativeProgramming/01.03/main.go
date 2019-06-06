// "Loops and Branches"
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	myNum := 37

	for {
		compNum := rand.Intn(100)
		if myNum > compNum {
			fmt.Println(compNum, "is too small")
		} else if myNum < compNum {
			fmt.Println(compNum, "is too big")
		} else {
			fmt.Println(compNum, "was the number!")
			break
		}
	}
}
