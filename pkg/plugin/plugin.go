// Copyright (C) 2019 Vanessa Sochat.
// This Source Code Form is subject to the terms of the
// Mozilla Public License, v. 2.0. If a copy of the MPL was not distributed
// with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

package plugin
 

// newFilterPlugin returns a filter plugin
func NewFilterPlugin(name string, usage string) *FilterPlugin {
	var functions = PluginFunctions{}
	var plugin = &FilterPlugin{Func: functions}
	plugin.configure(name, usage)
	return plugin
}

// newSinkPlugin returns a sink plugin
//func newSinkPlugin(name string, usage string) *SinkPlugin {
//	var functions = &PluginFunctions{}
//	var plugin = &SinkPlugin{Func: functions}
//	plugin.configure(name, usage)
//	return plugin
//}
