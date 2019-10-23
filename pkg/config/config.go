// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config
 
import (
	"encoding/json"
	"fmt"
)


// SyntaxShape is the kind of argument, and since it's easiest in Go to pass
// everything as a String, I suggest having everything returned as String
// and then doing custom parsing on returned values from config.GetParams the
// types here are only used to limit scope of what the user can provide. See
// https://github.com/nushell/nushell/blob/master/src/parser/hir/syntax_shape.rs#L45
type SyntaxShape string
const (
	ShapeString SyntaxShape = "String"
	ShapeInt SyntaxShape = "Int"
	ShapePath SyntaxShape = "Path"
	ShapeNumber SyntaxShape = "Number"
)

// ArgType are limited to optional, Switch (boolean) or Mandatory (required)
type ArgType string
const (
    Optional ArgType = "Optional"
    Switch ArgType = "Switch"
    Mandatory ArgType = "Mandatory"
)

// NamedParam is used as an entry in NamedParams
type NamedParam map[ArgType]SyntaxShape

// NamedParams are general string[string] map for the config
type NamedParams map[string]NamedParam

// Config holds configuration information for the plugin
type Config struct {
	Name	string			`json:"name"`
	Usage	string			`json:"usage"`
	Positional	[]string	`json:"positional"`
	RestPositional []string		`json:"rest_positional"`
	Named	NamedParams		`json:"named"`
	IsFilter	bool		`json:"is_filter"`
}

// ConfigResponse is specifically to return Jsonrpc with Params.ConfigResponseParams
type ConfigResponse struct {
	Jsonrpc string			`json:"jsonrpc"`
	Method string			`json:"method"`
	Params *ConfigResponseParams	`json:"params"`
}

// ConfigResponseParams add another level of nesting for {"Ok": Config}
type ConfigResponseParams map[string]Config


// AddNamedParam will take a key, argument type and shape and add to config.Named
func (config Config) AddNamedParam(key string, argType ArgType, syntaxShape SyntaxShape) {
	newParam := NamedParam{}
	newParam[argType] = syntaxShape
	config.Named[key] = newParam
}

// printConfigResponse will print the config json response to the terminal.
// generates an ConfigResponse with Params.ConfigResponseParams
func (config Config) PrintConfigResponse() error {

	responseParams := &ConfigResponseParams{"Ok": config}
	
	// Wrap params in json response
	response := &ConfigResponse{
		Jsonrpc: "2.0",
		Method: "response",
		Params: responseParams}

	// Serialize the struct to json, exit out if there is an error
	jsonString, err := json.Marshal(response) 
	if err != nil {
		return err
	}

	// Write the response to stdout
	fmt.Println(string(jsonString))
	return nil
}
