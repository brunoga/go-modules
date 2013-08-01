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
)

// GenericModule defines a generic module implementation. This is mainly to be
// used for struct embedding as it is not very useful by itself.
type GenericModule struct {
	name       string
	version    string
	genericId  string
	specificId string
	moduleType string
	parameters *ParameterMap
	ready      bool
}

// NewGenericModule creates and returns a new GenericModule.
func NewGenericModule(name, version, genericId, specificId, moduleType string) *GenericModule {
	return &GenericModule{
		name:       name,
		version:    version,
		genericId:  genericId,
		specificId: specificId,
		moduleType: moduleType,
		parameters: &ParameterMap{},
		ready:      false,
	}
}

func (gm *GenericModule) Name() string {
	return gm.name
}

func (gm *GenericModule) Version() string {
	return gm.version
}

func (gm *GenericModule) GenericId() string {
	return gm.genericId
}

func (gm *GenericModule) SpecificId() string {
	return gm.specificId
}

func (gm *GenericModule) Type() string {
	return gm.moduleType
}

func (gm *GenericModule) Register() error {
	return nil
}

func (gm *GenericModule) Unregister() error {
	return nil
}

func (gm *GenericModule) Parameters() *ParameterMap {
	return gm.parameters
}

func (gm *GenericModule) Configure(parameters *ParameterMap) error {
	gm.parameters = parameters
	return nil
}

func (gm *GenericModule) Duplicate(specificId string) (Module, error) {
	return nil, fmt.Errorf("generic module can not be duplicated")
}

func (gm *GenericModule) Ready() bool {
	return gm.ready
}

// SetReady allows setting the ready status of the module.
func (gm *GenericModule) SetReady(ready bool) {
	gm.ready = ready
}
