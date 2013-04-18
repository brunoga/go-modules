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

// Package modules implements support for generic modules. It allows programs to
// easily support the concept of having extendable code by providing an
// interface for automatic registration querying and usage of existing modules.
// Each module is completelly self contained and, usually, creating a module is
// just a matter of adding a new source code file for it without having to
// change anything else in the code that would use it.
package modules

import (
	"fmt"
)

// ParameterMap defines a list of parameters used for configuring a module.
type ParameterMap map[string]string

// Module is the interface that all modules must adhere to. This includes
// methods required by all module implementations. Note this is the bare
// minimum required and modules that just implement this are pretty useless.
type Module interface {
	// Name returns the human-readable name for the module.
	Name() string

	// Version returns the human-readable version for the module.
	Version() string

	// Id returns a short identifier for this module. This must be unique
	// among all registered modules and is enforced during module
	// registration.
	Id() string

	// Type returns the specific module type. Programs will usually only
	// have modules of one specific type, but not necessarily.
	Type() string

	// Register does any initialization required during module
	// registration. It is implicitly called by RegisterModule.
	Register() error

	// Unregister does any cleanup required during module removal. It is
	// implicitly called by UnregisterModule.
	Unregister() error

	// Parameters returns a ParameterMap with the configurable settings for
	// this module as keys and empty strings as values (except for entries
	// with default values). One could, then, fill/change the values in this
	// ParameterMap and pass it to Configure.
	Parameters() *ParameterMap

	// Configure configures the module usin the given ParameterMap. Usually
	// this needs to be called before actually trying to use the module.
	// Parameters above can be used to check what are the possible
	// configuration options.
	Configure(params *ParameterMap) error

	// iReady returns true if the module is ready to be used. False
	// otherwise. If the module requires configuration, this will return
	// false until Configure is called succesfully.
	Ready() bool
}

// ModuleMap is a container for a list of modules keyed by id.
type ModuleMap map[string]Module

// registeredModulesByType is a per type container for all registered modules.
var registeredModulesByType map[string]ModuleMap

// registeredModulesById is a per id container for all reistered modules.
var registeredModulesById ModuleMap

func init() {
	registeredModulesByType = make(map[string]ModuleMap)
	registeredModulesById = make(ModuleMap)
}

// RegisterModule registers a new module for usage.
func RegisterModule(module Module) error {
	moduleId := module.Id()
	_, ok := registeredModulesById[moduleId]
	if ok {
		return fmt.Errorf("id colision detected : %q", moduleId)
	}

	err := module.Register()
	if err != nil {
		return err
	}

	moduleType := module.Type()

	if registeredModulesByType[moduleType] == nil {
		registeredModulesByType[moduleType] = make(ModuleMap)
	}

	registeredModulesByType[moduleType][moduleId] = module
	registeredModulesById[moduleId] = module

	return nil
}

// UnregisterModule unregisters the given module.
func UnregisterModule(module Module) error {
	moduleType := module.Type()
	typeModuleList := registeredModulesByType[moduleType]
	if len(typeModuleList) != 0 {
		for id, m := range typeModuleList {
			if m == module {
				delete(typeModuleList, id)
				if len(registeredModulesByType[moduleType]) == 0 {
					delete(registeredModulesByType, moduleType)
				}

				moduleId := module.Id()
				delete(registeredModulesById, moduleId)

				return nil
			}
		}
	}

	return fmt.Errorf("module not found")
}

// GetModulesByType returns a ModuleMap with all modules of the iven type.
func GetModulesByType(moduleType string) ModuleMap {
	return registeredModulesByType[moduleType]
}

// GetModuleById returns the module represented by the given moduleId.
func GetModuleById(moduleId string) Module {
	return registeredModulesById[moduleId]
}

// GetAllModules returns a ModuleMap containing all registered modules.
func GetAllModules() ModuleMap {
	return registeredModulesById
}
