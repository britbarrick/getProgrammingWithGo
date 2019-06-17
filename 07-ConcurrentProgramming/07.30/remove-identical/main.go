// Goroutines and Concurrency -- remove-identical
package main

import "fmt"

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"a", "b", "b", "c", "d", "d", "d"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	last := ""
	for item := range upstream {
		if last != item {
			downstream <- item
		}
		last = item
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}

}
