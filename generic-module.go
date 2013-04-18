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

// GenericModule defines a generic module implementation. This is mainly to be
// used for struct embedding as it is not very useful by itself.
type GenericModule struct {
	name string
	version string
	id string
	moduleType string
	parameters *ParameterMap
	ready bool
}

// NewGenericModule creates and returns a new GenericModule.
func NewGenericModule(name, version, id, moduleType string) *GenericModule {
	return &GenericModule{
		name: name,
		version: version,
		id: id,
		moduleType: moduleType,
		parameters: &ParameterMap{},
		ready: false,
	}
}

func (gm *GenericModule) Name() string {
	return gm.name
}

func (gm *GenericModule) Version() string {
	return gm.version
}

func (gm *GenericModule) Id() string {
	return gm.id
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

func (gm *GenericModule) Ready() bool {
	return gm.ready
}

// SetReady allows setting the ready status of the module.
func (gm *GenericModule) SetReady(ready bool) {
	gm.ready = ready
}
