// Copyright 2013 Bruno Albuquerque (bga@bug-br.org.br).
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package main

import (
	base_modules "github.com/brunoga/go-modules"
)

var ModuleApiVersion = 1

type SimplePluginModule struct {
	*base_modules.GenericModule
}

func init() {
	base_modules.RegisterModule(&SimplePluginModule{
		base_modules.NewGenericModule("Simple Plugin Module", "1.0.0",
			"simple-plugin-module", "", "sample-module"),
	})
}
