// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main
 
import nu "github.com/vsoch/nu-plugin/pkg/plugin"


// sink is required for you to implement to parse parameters and print
// whatever you like to stdout!  It should take the plugin object passed, 
// along with params from nushell. You are responsible 
// for parsing the params, and can use plugin.Func.GetNamedParams(params)
func sink(plugin *nu.SinkPlugin, params interface{}) {

	// a map[string]interface{} with keys, values
	namedParams := plugin.Func.GetNamedParams(params)

	var color string

	// First pass we look for a custom color
	for name, value := range namedParams {
		switch name { 
			case "color":
				color = value.(string)
		}
	}

	// Second pass, look for fork and/or spoon
	for name, _ := range namedParams {

		if name == "fork" {
			printFork(color)
		} else if name == "spoon" {
			printSpoon(color)
		}
	}
}

func main() {

	name := "salad"
	usage := "Make a punny salad, with fork or spoon."
	plugin := nu.NewSinkPlugin(name, usage)

	// Add named arguments: key, config.SyntaxShape, config.ArgType
	// Switch should be provided with empty string
	plugin.Config.AddNamedParam("fork", "Switch", "")
	plugin.Config.AddNamedParam("spoon", "Switch", "")
	plugin.Config.AddNamedParam("color", "Optional", "String")

	// Run the filter function
	plugin.Run(sink)
}
