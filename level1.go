package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// colors returns a list of 25 colors. The colors can be any string representing
// a CSS color.
// Some of the possible formats: "blue", "rgb(42, 42, 42)", "#FF00FF".
//
// Step 1: return an array of 25x the same color (learn about loops and arrays)
// Step 2: alternate between 2 colors (learn about conditions)
// Step 3: make it a gradient (learn to build strings)
func colors() []string {
	// Create an array to hold the 25 colors
	colorArray := make([]string, 25)

	// Each step size for the gradient
	step := 255 / 24 // 255 is decimal for FF, and 24 is 25 - 1 (since we start from 0)

	for i := 0; i < 25; i++ {
		// Decrease the red component by step for each iteration
		redComponent := 255 - (step * i)

		// Increase the green component by step for each iteration
		greenComponent := step * i

		// Create the color string
		colorArray[i] = fmt.Sprintf("#%02X%02XFF", redComponent, greenComponent)
	}

	return colorArray
}

// No need to edit below this line

type colorsResponse struct {
	Colors []string `json:"colors"`
}

func Level1Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(colorsResponse{
		Colors: colors(),
	})
}
