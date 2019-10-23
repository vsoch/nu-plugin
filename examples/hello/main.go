// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main
 
import (
	"fmt"
	nu "github.com/vsoch/nu-plugin/pkg/plugin"
)

// sink is required for you to implement to parse parameters and print
// whatever you like to stdout!  It should take the plugin object passed, 
// along with params from nushell. You are responsible 
// for parsing the params, and can use plugin.Func.GetNamedParams(params)
func sink(plugin *nu.SinkPlugin, params interface{}) {

	// a map[string]interface{} with keys, values
	namedParams := plugin.Func.GetNamedParams(params)

	message := "Hello"
	excited := false

	// First pass we look for booleans
	for name, _ := range namedParams {
		switch name { 
			case "excited":
				excited = true
		}
	}

	// Second pass, look for string args
	for name, value := range namedParams {

		if name == "name" {
			message = message + " " + value.(string)
		}
	}

	// Add an exclamation point?
	if excited {
		message = message + "!"
	}

	fmt.Println(message)
}

func main() {

	name := "hello"
	usage := "A friendly plugin"
	plugin := nu.NewSinkPlugin(name, usage)

	// Add named arguments: key, config.SyntaxShape, config.ArgType
	// Switch should be provided with empty string
	plugin.Config.AddNamedParam("excited", "Switch", "")
	plugin.Config.AddNamedParam("name", "Optional", "String")

	// Run the filter function
	plugin.Run(sink)
}
