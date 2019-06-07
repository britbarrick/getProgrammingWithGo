// First-class Functions
package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var off kelvin = 12
	sensor := calibrate(realSensor, off)
	fmt.Println(sensor())
	off++
	fmt.Println(sensor())

	sensor1 := calibrate(fakeSensor, 5)
	fmt.Println(sensor1())
	fmt.Println(sensor1())
	fmt.Println(sensor1())
	fmt.Println(sensor1())

}
