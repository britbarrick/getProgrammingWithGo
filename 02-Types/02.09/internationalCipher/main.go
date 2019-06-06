// Multilingual Text
package main

import "fmt"

func main() {
	message := "Hola Estación Espacial Internacional"

	for i := range message {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
}
