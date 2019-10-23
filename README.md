# Nushell Plugin in GoLang

This is a base library for generating a nushell plugin in Golang.
As nushell is under development, it's likely that not all features are implemented
here! Please [open an issue](https://www.github.com/vsoch/nu-plugin/issues)
if you need help.

## Filter Plugin

I've done a basic [filter example](examples/len) to calculate the length of a string, 
and the usage is summarized as follows:


```go
package main
 
import nu "github.com/vsoch/nu-plugin/pkg/plugin"


// filter will read stream from nushell and print a response
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
```

For more details, see the [full example](examples/len).

## Sink Plugin

We have two examples for Sink plugins:

 - [examples/salad](examples/salad) prints colored salad puns to the screen
 - [examples/hello](examples/hello) is a more basic example to say hello

