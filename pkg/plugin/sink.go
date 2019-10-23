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

// SinkPlugin represents an interface for a Nushell Sink Plugin. 
// It includes a configuration, along with supporting functions
type SinkPlugin struct {
	Config	configlib.Config
	Func	PluginFunctions
}

// configure the sink plugin
// TODO need to add Named and Positional to plugin
func (plugin *SinkPlugin) configure(name string, usage string) {
	var config = configlib.Config{
		Name: name,
		Usage: usage,
		Named: configlib.NamedParams{},
		Positional: make([]string, 0),
		IsFilter: false}
	plugin.Config = config
}

// getLogfile returns the name of the logfile in /tmp
func (plugin *SinkPlugin) getLogfile() string {
	logfile := path.Join(os.TempDir(), "nu_plugin_" + plugin.Config.Name + ".log")
	return logfile
}


// printConfigResponse is a wrapper to Config.printConfigResponse
func (plugin *SinkPlugin) printConfigResponse() {
	plugin.Config.PrintConfigResponse()
}

// printConfigResponse is a wrapper to Config.printConfigResponse
func (plugin *SinkPlugin) getExecutableName() string {
	return "nu_plugin_" + plugin.Config.Name
}


// Run the sink, take the user's function as a parameter along with the plugin
func (plugin *SinkPlugin) Run(sinkFunc func(plugin *SinkPlugin, stringValue interface{})) {

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

			} else if method == "sink" {
				logger.Println("Request for sink", line)
				if params, ok := line["params"]; ok {
					sinkFunc(plugin, params)
				}
				break
			}

		} else {
			break
		}
	}
}
