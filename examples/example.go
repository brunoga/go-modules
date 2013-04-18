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

	base_modules "bga/modules"
	example_modules "bga/modules/examples/modules"
)

func main() {
	fmt.Println("Number of sample modules :",
		len(base_modules.GetModulesByType("sample-module")))
	fmt.Println("Total number of modules  :",
		len(base_modules.GetAllModules()))

	simpleModule := base_modules.GetModuleById("simple-module").
		(*example_modules.SimpleModule)
	if (simpleModule != nil) {
		fmt.Println(simpleModule.Id(), simpleModule.Name(),
			simpleModule.Version())
	}

	completeModule := base_modules.GetModuleById("complete-module").
		(*example_modules.CompleteModule)
	if completeModule != nil {
		fmt.Println(completeModule.Id(), completeModule.Name(),
                        completeModule.Version())
		err := completeModule.DoWork()
		if err != nil {
			fmt.Printf("error doing work : %v\n", err)
		}

		parameters := completeModule.Parameters()
		(*parameters)["text1"] = "This is text1"
		completeModule.Configure(parameters)

		err = completeModule.DoWork()
		if err != nil {
			fmt.Printf("error doing work : %v\n", err)
		}
	}
}

