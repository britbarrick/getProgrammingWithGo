// A Little Structure
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Name string  `json:"landing site"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations := []location{
		{"Bradbury Landing", -4.5895, 137.4417},
		{"Columbia Memorial Station", -14.5684, 175.472636},
		{"Challenger Memorial Station", -1.9462, 354.4734},
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
