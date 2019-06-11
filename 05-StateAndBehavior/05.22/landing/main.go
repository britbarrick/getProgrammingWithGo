// Go's Got No Class -- Landing
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string  `json:"Name"`
	Lat  float64 `json:"Latitude"`
	Long float64 `json:"Longitude"`
}

func main() {
	locations := []location{
		{"Columbia Memorial Station", -14.5683, 175.4725},
		{"Challenger Memorial Station", -1.9461, 354.4733},
		{"Bradbury Landing", -4.5895, 137.4416},
		{"Elysium Planitia", 4.5000, 135.9000},
	}

	bytes, err := json.MarshalIndent(locations, "", " ")
	exitOnError(err)

	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
