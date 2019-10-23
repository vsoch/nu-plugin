// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package plugin

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	path "path/filepath"

	configlib "github.com/vsoch/nu-plugin/pkg/config"
)

// FilterPlugin represents an interface for a Nushell Filter Plugin. 
// It includes a configuration, along with supporting functions
type FilterPlugin struct {
	Config	configlib.Config
	Func	PluginFunctions
}

// configure is called upon plugin generation
// TODO need to add Named and Positional to plugin
func (plugin *FilterPlugin) configure(name string, usage string) {
	var config = configlib.Config{
		Name: name,
		Usage: usage,
		Named: configlib.NamedParams{},
		Positional: make([]string, 0),
		IsFilter: true}
	plugin.Config = config
}

// getLogfile returns the name of the logfile in /tmp
func (plugin *FilterPlugin) getLogfile() string {
	logfile := path.Join(os.TempDir(), "nu_plugin_" + plugin.Config.Name + ".log")
	return logfile
}


// printConfigResponse is a wrapper to Config.printConfigResponse
func (plugin *FilterPlugin) printConfigResponse() {
	plugin.Config.PrintConfigResponse()
}

// printConfigResponse is a wrapper to Config.printConfigResponse
func (plugin *FilterPlugin) getExecutableName() string {
	return "nu_plugin_" + plugin.Config.Name
}

// startFilter is a wrapper to plugin.printEmptyResponse
func (plugin *FilterPlugin) beginFilter() error {
	return plugin.printEmptyResponse()
}

// endFilter is a wrapper to plugin.printEmptyResponse
func (plugin *FilterPlugin) endFilter() error {
	return plugin.printEmptyResponse()
}

// printEmptyResponse will print an EmptyResponse.
// the intende use case is for an end_filter or start_filter
// generates an ArrayResponse with Params.EmptyResponseParams
func (plugin *FilterPlugin) printEmptyResponse() error {

	emptyArray := make([]string, 0)
	responseParams := &EmptyResponseParams{"Ok": emptyArray}
	
	// Wrap params in json response
	response := &EmptyResponse{
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

// Run the filter, should be called in main of the implemented plugin
func (plugin *FilterPlugin) Run(filterFunc func(plugin *FilterPlugin, stringValue interface{})) {


	// Set up temporary logger
	f, err := os.OpenFile(plugin.getLogfile(),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()	
	logger := log.New(f, plugin.getExecutableName(), log.LstdFlags)

	// Read into json decoded
	decoder := json.NewDecoder(os.Stdin)

	line := make(map[string]interface{})

	for {
		err := decoder.Decode(&line) 
		if err != nil {
			fmt.Errorf("unable to read json: %s", err)
		} 

		// look for a method in the line
		if method, ok := line["method"]; ok {	
	
			// Case 1: method is config
			if method == "config" {
				logger.Println("Request for config", line)
			        plugin.printConfigResponse()
				break

			} else if method == "begin_filter" {
				logger.Println("Request for begin filter", line)
				plugin.beginFilter()

			} else if method == "filter" {
				logger.Println("Request for filter", line)
				if params, ok := line["params"]; ok {
					filterFunc(plugin, params)
				}

			} else if method == "end_filter" {
				logger.Println("Request for end filter", line)
				plugin.endFilter()
				break
			}

		} else {
			break
		}
	}
}
