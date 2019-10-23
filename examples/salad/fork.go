// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"os"
)

// selectForkPun randomly from the list
func selectForkPun() string {

	// Create a "slice" (array) of puns
	puns := []string{" You're done!",
		" Cut it out.",
		" Fork off, already.",
		" You want a peas of me?",
		" Have a knife day? Fork you!",
		" My life purpose: I cut butter.",
		" Forkin' repos, that's what I do.",
		" What the F*rk are you asking me for?!",
		" You think you have problems? I'm a fork.",
		" Take the one less traveled, they said...",
		" In Go an array is a slice. Utensil discrimination!",
		" I'm not much of a traveler, but people take me anyway.",
		" I can't help with yo' Momma, I'm not that kind of fork."}

	// Randomly select one
	return selectRandom(puns...)

}

// selectFork from the ascii options
func selectFork() string {

	forks := make([]string, 5)

	// source
	// ascii.co.uk/art/fork

	forks[0] = `

                   ________  .====
                  [________>< :===
                             '====`

	forks[1] = `

          _________________  .========
         [_________________>< :======
                             '========`

	forks[2] = `

                            _
                           / )
                     |||| / /
                     ||||/ /
                     \__(_/
                      ||//
                      ||/
                      ||
                     (||
                      ""`

	forks[3] = `

                       /\
                      //\\
                     //  \\
                 ^   \\  //   ^
                / \   )  (   / \ 
                ) (   )  (   ) (
                \  \_/ /\ \_/  /
                 \__  _)(_  __/
                    \ \  / /
                     ) \/ (
                     | /\ |
                     | )( |
                     | )( |
                     | \/ |
                     )____(
                    /      \
                    \______/`

	forks[4] = `

                                   ⎯⎯∈

               `

	// Randomly select one
	return selectRandom(forks...)
}

// printFork to the terminal
func printFork(color string) {

	pun := selectForkPun()
	fork := selectFork()
	color = selectColor(color)

	fmt.Println()
	fmt.Println(pun, color, fork, "\033[0m") // off sequence to end color
	fmt.Println()

	os.Exit(0)

}
