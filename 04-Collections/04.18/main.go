// A Bigger Slice
package main

import "fmt"

type number []int

func main() {
	var num number

	for i := 0; i < 20; i++ {
		num = append(num, i)
		fmt.Println(cap(num), num)
	}
}
