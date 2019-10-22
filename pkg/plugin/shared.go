// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package plugin
 
import (
	"encoding/json"
	"fmt"
)

// NushellFunctions has shared functions for a plugin
type PluginFunctions struct {}


// getStringPrimitive returns stringValue["item"]["Primitive"]["String"]
func (caller *PluginFunctions) GetStringPrimitive(stringValue interface{}) string {
	
	// I hope there is a more elegant way to do this
	jsonValues := stringValue.(map[string]interface{})
	item := jsonValues["item"].(map[string]interface{})
	primitive := item["Primitive"].(map[string]interface{})
	finalString := primitive["String"].(string)
	return finalString
}


// getIntPrimitive returns stringValue["item"]["Primitive"]["Int"]
func (caller *PluginFunctions) GetIntPrimitive(stringValue interface{}) int {
	
	// I hope there is a more elegant way to do this
	jsonValues := stringValue.(map[string]interface{})
	item := jsonValues["item"].(map[string]interface{})
	primitive := item["Primitive"].(map[string]interface{})
	value := primitive["Int"].(int)
	return value
}

// printIntResponse will print a json response to the terminal. The 
// generates a JsonResponse with Params.FinalResponseParams
func (caller *PluginFunctions) PrintIntResponse(value int, tag Tag) error {

	intResponse := IntResponse{}
	intResponse.Item.Primitive.Int = value
	intResponse.Tag = tag

	nestedParams := IntResponseParams{"Value": intResponse}
	params := IntResponseWrapper{"Ok": nestedParams}

	arrayResponse := IntArrayResponse{}
	arrayResponse = append(arrayResponse, params)

	finalResponseParams := FinalIntResponseParams{"Ok": arrayResponse}
	
	// Wrap params in json response
	response := &JsonIntResponse{
		Jsonrpc: "2.0",
		Method: "response",
		Params: finalResponseParams}

	// Serialize the struct to json, exit out if there is an error
	jsonString, err := json.Marshal(response) 
	if err != nil {
		return err
	}

	// Write the response to stdout
	fmt.Println(string(jsonString))
	return nil
}


// printStringResponse will print a json response to the terminal. The 
// generates a JsonResponse with Params.FinalResponseParams
func (caller *PluginFunctions) PrintStringResponse(value string, tag Tag) error {

	stringResponse := StringResponse{}
	stringResponse.Item.Primitive.String = value
	stringResponse.Tag = tag

	nestedParams := StringResponseParams{"Value": stringResponse}
	params := StringResponseWrapper{"Ok": nestedParams}

	arrayResponse := StringArrayResponse{}
	arrayResponse = append(arrayResponse, params)

	finalResponseParams := FinalStringResponseParams{"Ok": arrayResponse}
	
	// Wrap params in json response
	response := &JsonStringResponse{
		Jsonrpc: "2.0",
		Method: "response",
		Params: finalResponseParams}

	// Serialize the struct to json, exit out if there is an error
	jsonString, err := json.Marshal(response) 
	if err != nil {
		return err
	}

	// Write the response to stdout
	fmt.Println(string(jsonString))
	return nil
}
