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

package modules

import (
	"fmt"

	base_modules "bga/modules"
)

type CompleteModuleInterface interface {
	base_modules.Module
	DoWork() error
}

type CompleteModule struct {
	*base_modules.GenericModule
	text1 string
	text2 string
}

func (cm *CompleteModule) Parameters() *base_modules.ParameterMap {
	return &base_modules.ParameterMap{
		"text1" : "",
		"text2" : "Default text",
	}
}

func (cm *CompleteModule) Configure(parameters *base_modules.ParameterMap) error {
	text1Parameter, ok := (*parameters)["text1"]
	if !ok {
		return fmt.Errorf("required parameter text1 is missing")
	}

	text2Parameter := (*parameters)["text2"]

	cm.text1 = text1Parameter
	cm.text2 = text2Parameter

	cm.SetReady(true)

	return nil
}

func (cm *CompleteModule) DoWork() error {
	if !cm.Ready() {
		return fmt.Errorf("module not ready")
	}

	fmt.Println("text1 :", cm.text1)
	fmt.Println("text2 :", cm.text2)

	return nil
}

func registerCompleteModule(module CompleteModuleInterface) error {
	return base_modules.RegisterModule(module)
}

func init() {
	registerCompleteModule(&CompleteModule{
		base_modules.NewGenericModule("Complete Module", "1.0.0",
			"complete-module", "sample-module"),
		"",
		"",
	})
}

