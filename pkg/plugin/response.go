// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package plugin


// IntResponse is nested under Primitive -> Int -> value
type IntResponse struct {
	Item struct {
    		Primitive struct {
			Int int `json:"Int"`
		} `json:"Primitive"`
	} `json:"item"`
	Tag Tag	`json:"tag"`
}

// IntResponseParams support inner {"Value": IntResponse} with Item.Primitive.Int
type IntResponseParams map[string]IntResponse

// IntResponseWrapper support inner {"Ok": IntResponseParams}
type IntResponseWrapper map[string]IntResponseParams

// IntArrayResponseParams are an array of []IntResponseWrapper for JsonIntResponse.Params
type IntArrayResponse []IntResponseWrapper

// FinalIntResponseParams wraps IntArrayResponse: {"Ok": IntArrayResponse}
type FinalIntResponseParams map[string]IntArrayResponse

// IntJsonResponse returns FinalIntResponseParams
type JsonIntResponse struct {
	Jsonrpc string			`json:"jsonrpc"`
	Method string			`json:"method"`
	Params FinalIntResponseParams	`json:"params"`
}



// StringResponse is nested under Primitive -> String -> value
type StringResponse struct {
	Item struct {
    		Primitive struct {
			String string `json:"String"`
		} `json:"Primitive"`
	} `json:"item"`
	Tag Tag	`json:"tag"`
}

// StringResponseParams support inner {"Value": StringResponse} with Item.Primitive.String
type StringResponseParams map[string]StringResponse

// IntResponseWrapper support inner {"Ok": IntResponseParams}
type StringResponseWrapper map[string]StringResponseParams

// ArrayResponseParams are an array of []ResponseParams for JsonResponse.Params
type StringArrayResponse []StringResponseWrapper

// FinalIntResponseParams wraps IntArrayResponse: {"Ok": IntArrayResponse}
type FinalStringResponseParams map[string]StringArrayResponse

// IntJsonResponse returns FinalIntResponseParams
type JsonStringResponse struct {
	Jsonrpc string				`json:"jsonrpc"`
	Method string				`json:"method"`
	Params FinalStringResponseParams	`json:"params"`
}



// ArrayParams are intended for end_filter and begin_filter
type ArrayParams []string

// EmptyResponseParams supports {"Ok": ArrayParams} (e.g., {"Ok": []})
type EmptyResponseParams map[string]ArrayParams

// EmptyResponse returns EmptyResponseParams
type EmptyResponse struct {
	Jsonrpc string			`json:"jsonrpc"`
	Method string			`json:"method"`
	Params *EmptyResponseParams	`json:"params"`
}
