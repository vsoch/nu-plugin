// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package plugin


// Tag is passed from stream during filter, we parse and pass forward
type Tag struct {
	Anchor interface{}	`json:"anchor"`
	Span map[string]int	`json:"span"`
}

// getTag from the filter input to return in response
func (caller *PluginFunctions) GetTag(stringValue interface{}) Tag {
	
	// I hope there is a more elegant way to do this
	jsonValues := stringValue.(map[string]interface{})

	// Create the span to hold a start and end, and empty tag
	span := map[string]int{}
	tag := Tag{}

	// Real nushell invokation will include a tag
	if tagGroup, ok := jsonValues["tag"].(map[string]interface{}); ok {

		spanGroup := tagGroup["span"].(map[string]interface{})

		// Create the span, it has a start and end
		span["start"] = int(spanGroup["start"].(float64))
		span["end"] = int(spanGroup["end"].(float64))
		tag.Span = span

		// If anchor isn't nil, add it (not sure if type is correct)
		if anchor, ok := tagGroup["anchor"].(interface{}); ok {
			tag.Anchor = anchor
		}

	// Otherwise generate a dummy one for local testing
	} else {
		span["start"] = 0
		span["end"] = 0
		tag.Span = span
	}

	return tag
}
