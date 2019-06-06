// Multilingual Text
package main

import "fmt"

func main() {
	quote := "L fdph, L vdz, L frqtxhuhg"

	for i := range quote {
		c := quote[i]
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
}
