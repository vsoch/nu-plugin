// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"math/rand"
	"time"
)



// selectRandom from strings
func selectRandom(choices ...string) string {

	// Seed random number generator, won't be random without it
	rand.Seed(time.Now().Unix())

	// Make random selection
	return choices[rand.Intn(len(choices))]
}


// selectRandomMap from a choice of maps
func selectRandomMap(choices map[string]string) string {

        // If given a map, convert to array of values first
        values := make([]string, len(choices))
        i := 0
        for k := range choices {
            values[i] = choices[k]
            i++
        } 

        return selectRandom(values...)
}


// selectColor out of standard terminal options
func selectColor(color string) string {

	colors := map[string]string{
		"purple":  "\033[95m",
		"yellow":  "\033[93m",
		"red":     "\033[91m",
		"darkred": "\033[31m",
		"cyan":    "\033[36m"}

	if color != "" {
		if _, ok := colors[color]; ok {
			return colors[color]
		}
	}

	return selectRandomMap(colors)
}
