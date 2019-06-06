// Converting Between Types
package main

import "fmt"

func main() {
	fmt.Println(trueFalse("false"))

}

func trueFalse(val string) bool {
	tf := true
	switch val {
	case "true", "yes", "1":
		tf = true
	case "false", "no", "0":
		tf = false
	}
	return tf
}
