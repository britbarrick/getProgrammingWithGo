// Big Numbers
package main

import (
	"fmt"
)

func main() {
	const lightspeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365
	const canisDist = 236000000000000000

	ltYears := canisDist / lightspeed / secondsPerDay / daysPerYear
	fmt.Println("Canis is", ltYears, "lightyears from the sun.")
}
