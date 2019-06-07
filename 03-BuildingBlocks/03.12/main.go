package main

import "fmt"

// kelvinToCelsius converts K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// celsiusToFahrenheit converts °C to °F
func celsiusToFahrenheit(c float64) float64 {
	c = (c * (9.0 / 5.0)) + 32.0
	return c
}

// kelvinToFahrenheit converts K to °F
func kelvinToFahrenheit(k float64) float64 {
	k -= 459.67
	return k
}

func main() {
	kelvin := 0.0
	fmt.Printf("%vK is %v°C.\n", kelvin, kelvinToCelsius(kelvin))
	fmt.Printf("%vK is %v°F\n", kelvin, kelvinToFahrenheit(kelvin))
}
