// A Few Pointers
package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y++
}

func (t *turtle) down() {
	t.y--
}

func (t *turtle) right() {
	t.x++
}

func (t *turtle) left() {
	t.x--
}

func main() {
	turtle := turtle{
		x: 4,
		y: 5,
	}

	turtle.down()
	turtle.down()
	turtle.down()
	turtle.down()
	turtle.down()
	turtle.down()
	turtle.down()
	turtle.down()
	turtle.left()
	turtle.left()
	turtle.left()
	turtle.right()
	turtle.up()

	fmt.Println(turtle)

}
