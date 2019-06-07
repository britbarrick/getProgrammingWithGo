// Methods
package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// celsius converts K to °C
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

//kelvin converts °C to K
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32)
}

// celsius converts °F to °C
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func main() {
	var k kelvin
	var f fahrenheit = 32.0
	var c celsius
	fmt.Println(k.celsius(), k.fahrenheit(), f.celsius(), f.kelvin(), c.kelvin(), c.fahrenheit())
}
