// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package config
 
import (
	"encoding/json"
	"fmt"
)


// ConfigParams are general string[string] map for the config
type ConfigParams map[string]string

// Config holds configuration information for the plugin
type Config struct {
	Name	string			`json:"name"`
	Usage	string			`json:"usage"`
	Positional	[]string	`json:"positional"`
	RestPositional []string		`json:"rest_positional"`
	Named	ConfigParams		`json:"named"`
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
