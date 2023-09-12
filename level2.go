package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Level2 is about using structs and storing state.
// Use the struct below to store a state and make colors move!
//
// Step 1: make a square move at every invocation (learn about structs)
// Add a field to the `Level2` struct and update it in the `colors` function.
//
// Step 2: use defer to change the state (learn about defer).
// defer is an effective way of executing logic at the end of functions, try
// using it to update your state!
//
// Step 3: get the color as a constructor parameter (learn how to add parameters to a struct)
// Update the `NewLevel2` function to accept a color parameter.
type Level2 struct {
	state       int
	randomColor string
}

func NewLevel2(randomColor string) *Level2 {
	return &Level2{
		// FIXME: Initialize your value here
		state:       0,
		randomColor: randomColor,
	}
}

func (l *Level2) colors() []string {
	// Create an array to hold the 25 colors
	colorArray := make([]string, 25)

	defer func() {
		// This will update the state at the end of the function call
		l.state = (l.state + 1) % 25
	}()

	for i := range colorArray {
		if i == l.state {
			colorArray[i] = l.randomColor
		} else {
			step := 255 / 24
			redComponent := 255 - (step * i)
			greenComponent := step * i
			colorArray[i] = fmt.Sprintf("#%02X%02XFF", redComponent, greenComponent)
		}
	}

	return colorArray
}

// No need to edit below this line

func (l *Level2) Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(colorsResponse{
		Colors: l.colors(),
	})
}
