// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main
 
import nu "github.com/vsoch/nu-plugin/pkg/plugin"


// filter is required for you to implement to print the response needed
// It should take the plugin object passed, along with params from the 
// calling filterPlugin. It should run plugin.printStringResponse or 
// plugin.printIntResponse, depending on your response. You are responsible
// for parsing the params, and can use both plugin.Func.getStringPrimivite 
// and plugin.Func.getIntPrimivite
func filter(plugin *nu.FilterPlugin, params interface{}) {

	// can also be getIntPrimitive
	value := plugin.Func.GetStringPrimitive(params)

	// Put your logic here! In this case, we want a length
	intLength := len(value)

	// You must also return the tag with your response
	tag := plugin.Func.GetTag(params)

	// This can also be printStringResponse
	plugin.Func.PrintIntResponse(intLength, tag)

}


func main() {

	name := "len"
	usage := "Return the length of a string"
	plugin := nu.NewFilterPlugin(name, usage)

	// Run the filter function
	plugin.Run(filter)
}
