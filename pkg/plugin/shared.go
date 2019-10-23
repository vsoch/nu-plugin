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

// GetNamedParams, or a map of names to values provided from the params
// key. We pass as a stringvalue interface that is expected to look like
// 	[{'args': {'positional': None,
//	   'named': {'switch': {'tag': {'anchor': None,
//	      'span': {'start': 58, 'end': 64}},
//	     'item': {'Primitive': {'Boolean': True}}},
//	    'mandatory': {'tag': {'anchor': None, 'span': {'start': 20, 'end': 32}},
//	     'item': {'Primitive': {'String': 'MANDATORYARG'}}},
//	    'optional': {'tag': {'anchor': None, 'span': {'start': 44, 'end': 55}},
//	     'item': {'Primitive': {'String': 'OPTIONALARG'}}}}},
//	  'name_tag': {'anchor': None, 'span': {'start': 0, 'end': 7}}},
//	 []]
func (caller *PluginFunctions) GetNamedParams(stringValue interface{}) map[string]string {

	response := make(map[string]string)
	fmt.Println(stringValue)
	return response
}
//    # Just grab the args dictionary
//    input_params = input_params[0]
//    positional = input_params['args'].get('positional', [])
//    named = input_params['args'].get('named', {})

//    # We will return lookup dictionary of params
//    params = {}

//    # Keep a simple dictionary with values we know types for
//    for name, values in named.items():

//        # is it a String? Boolean?
//        value_type = list(values['item']['Primitive'].keys())[0]

//        if value_type == "String":
//            params[name] = values['item']['Primitive']['String']

//        elif value_type == "Boolean":
//            params[name] = values['item']['Primitive']['Boolean']

//        # If you use other types, add them here

//        else:
//            logging.info("Invalid paramater type %s:%s" %(name, values))

//    return params        


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
