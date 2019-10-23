// Copyright (C) 2018-2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"os"
)


// selectSpoonPun from the list
func selectSpoonPun() string {

	// Create a "slice" (array) of puns
	puns := []string{" See you spoon!",
		" They say I'm rather spoontaneous!",
		" Did you hear about... oh, too spoon?",
		" They call me the cereal kiler.",
		" Wanna spoon?",
		" My favorite actress? Reese Witherspoon!",
		" *singing* The dark side of... the spooooon!",
		" I'm a spoon. How would I know?",
		" If you want to be sharp, I'm the wrong utensil.",
		" I don't have a point. Go talk to Fork."}

	// Randomly select one
	return selectRandom(puns...)
}

// selectSpoon ascii from the list
func selectSpoon() string {

	spoons := make([]string, 5)

	// source
	// spoon.ascii.uk/

	spoons[0] = `

           ________   .==.
          [________>c((_  )
                      '=='

         `

	spoons[1] = `


        __________        .-"""-.
       /          ''''---' .'    \
       \__________....---. '.    /
                          '-...-'
        `

	spoons[2] = `

                              _
                             / \
          _..--"""""--.._    \_/
         /,_..-------.._,\    |
        |  ''''-----''''  |   |
         \               /   / \
          '.           .'    | |
            '--.....--'      \_/
        `

	spoons[3] = `

          ___           .-""-.
         /   '''---...-'.'  '\\
         \___...---"""-._-.__//
                         '---'

        `

	spoons[4] = `


           $$$$$$$$$
          $$$$$$$$$$$
         $$$$$$$$$$$$$$
        $$$$$$$$$$$$$$$$
        $$$$$$$$$$$$$$$$
        $$$$$$$$$$$$$$$$
         $$$$$$$$$$$$$$
          $$$$$$$$$$$$
           $$$$$$$$$$
            $$$$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
            $$$$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
             $$$$$$
              $$$$

        `

	// Randomly select one
	return selectRandom(spoons...)
}

// printSpoon to the terminal
func printSpoon(color string) {

	pun := selectSpoonPun()
	spoon := selectSpoon()
	color = selectColor(color)

	fmt.Println()
	fmt.Println(pun, color, spoon, "\033[0m") // off sequence to end color
	fmt.Println()

	os.Exit(0)

}
