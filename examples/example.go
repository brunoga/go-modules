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
	"fmt"

	base_modules "gopkg.in/brunoga/go-modules.v1"
	example_modules "gopkg.in/brunoga/go-modules.v1/examples/modules"
)

func main() {
	fmt.Println("Number of sample modules :",
		base_modules.GetModuleCountByType("sample-module"))
	fmt.Println("Total number of modules  :",
		base_modules.GetAllModulesCount())

	simpleModule := base_modules.GetModuleById("simple-module", "").(*example_modules.SimpleModule)
	if simpleModule != nil {
		fmt.Println(simpleModule.GenericId(), simpleModule.Name(),
			simpleModule.Version())
	}

	completeModule := base_modules.GetModuleById("complete-module", "").(*example_modules.CompleteModule)
	if completeModule != nil {
		fmt.Println(completeModule.GenericId(), completeModule.Name(),
			completeModule.Version())
		// This will fail as module is not configured.
		err := completeModule.DoWork()
		if err != nil {
			fmt.Printf("error doing work : %v\n", err)
		}

		parameters := completeModule.Parameters()
		(*parameters)["text1"] = "This is text1"
		completeModule.Configure(parameters)

		// And now itwill work as expected.
		err = completeModule.DoWork()
		if err != nil {
			fmt.Printf("error doing work : %v\n", err)
		}
	}

	// Add a new complete module with a different specific id.
	base_modules.RegisterModule(base_modules.NewGenericModule(
		"Complete Module / Specific", "1.0.0",
		"complete-module", "specific", "sample-module"))

	fmt.Println("After adding new module.")
	fmt.Println("Number of sample modules :",
		base_modules.GetModuleCountByType("sample-module"))
	fmt.Println("Total number of modules  :",
		base_modules.GetAllModulesCount())

	completeModuleSpecific := base_modules.GetModuleById("complete-module", "specific")
	if completeModuleSpecific == nil {
		fmt.Println("Could not get reference to specific module.")
	}
	completeModuleDefault := base_modules.GetDefaultModuleByGenericId("complete-module")
	if completeModuleDefault == nil {
		fmt.Println("Could not get reference to default module.")
	}

	fmt.Println(completeModuleSpecific.Name())
	fmt.Println(completeModuleDefault.Name())
}
