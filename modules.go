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

	// GenericId returns a short generic identifier for this module. Usually
	// it identifies the "class" of this Module (for example,
	// "directory-reader" and, together with SpecificId() below, must
	// uniquely identify an instance of this Module.
	GenericId() string

	// SpecificId returns a short identifier for this module. It serves to
	// identify a specific Module instance in a given "class" (see
	// GenericId()). For example, assuming a GenericId() of
	// "directory-reader", this could return something that represents a
	// specific directory. For example, it could return "home" to indicate
	// that it is a "directory-reader" that operates in the "home"
	// directory.
	SpecificId() string

	// Type returns the specific module type. This is used do group modules
	// in specific domains so one can query about all modules on them. A
	// type could, for example, be a refrence to the program that uses
	// those modules (in other words, it could be the program name).
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

	// Duplicate creates and registers a new instance of this module with
	// the given specific id. Returns a reference to the new Module and a
	// nil error on success and a non-nil error on failure. Note that
	// this must be implemented by each Module that cares about being
	// duplicated and it is the responsibility of each implementation to
	// register (calling RegisterModule()) the duplicate.
	Duplicate(specificId string) (Module, error)

	// Ready returns true if the module is ready to be used. False
	// otherwise. If the module requires configuration, this will return
	// false until Configure is called succesfully.
	Ready() bool
}

// ModuleMap is a container for Modules keyed by specific id.
type ModuleMap map[string]Module

// FullModuleMap is a container for ModuleMaps keyed by generic id.
type FullModuleMap map[string]ModuleMap

// registeredModulesByType is a per type container for all registered modules.
var registeredModulesByType map[string]FullModuleMap

// registeredModulesById is a per id container for all registered modules.
var registeredModulesById FullModuleMap

func init() {
	registeredModulesByType = make(map[string]FullModuleMap)
	registeredModulesById = make(FullModuleMap)
}

// RegisterModule registers a new module for usage.
func RegisterModule(module Module) error {
	genericModuleId := module.GenericId()
	specificModuleId := module.SpecificId()
	_, ok := registeredModulesById[genericModuleId][specificModuleId]
	if ok {
		return fmt.Errorf("id colision detected : %q / %q", genericModuleId, specificModuleId)
	}

	err := module.Register()
	if err != nil {
		return err
	}

	moduleType := module.Type()

	if registeredModulesByType[moduleType] == nil {
		registeredModulesByType[moduleType] = make(FullModuleMap)
	}
	if registeredModulesByType[moduleType][genericModuleId] == nil {
		registeredModulesByType[moduleType][genericModuleId] = make(ModuleMap)
	}
	if registeredModulesById[genericModuleId] == nil {
		registeredModulesById[genericModuleId] = make(ModuleMap)
	}

	registeredModulesByType[moduleType][genericModuleId][specificModuleId] = module
	registeredModulesById[genericModuleId][specificModuleId] = module

	return nil
}

// UnregisterModule unregisters the given module.
func UnregisterModule(module Module) error {
	moduleType := module.Type()
	genericModuleId := module.GenericId()
	specificModuleId := module.SpecificId()

	typeModuleMap := registeredModulesByType[moduleType][genericModuleId]
	for id, m := range typeModuleMap {
		if m == module {
			// Remove module from modules by type map.
			delete(typeModuleMap, id)
			if len(registeredModulesByType[moduleType][genericModuleId]) == 0 {
				// No more modules in generic id map. Delete it
				// too.
				delete(registeredModulesByType[moduleType], genericModuleId)
			}
			if len(registeredModulesByType[moduleType]) == 0 {
				// No more modules in this type. Delete it too.
				delete(registeredModulesByType, moduleType)
			}

			// Remove module from modules by id map.
			delete(registeredModulesById[genericModuleId], specificModuleId)
			if len(registeredModulesById[genericModuleId]) == 0 {
				// No more modules in generic id map. Delete it
				// too.
				delete(registeredModulesById, genericModuleId)
			}

			return nil
		}
	}

	return fmt.Errorf("module not found")
}

// GetModulesByType returns a FullModuleMap with all modules of the given type.
func GetModulesByType(moduleType string) FullModuleMap {
	return registeredModulesByType[moduleType]
}

// GetModuleCountByType returns the number of registered modules of a specific
// type.
func GetModuleCountByType(moduleType string) int {
	return countFullModuleMap(GetModulesByType(moduleType))
}

// GetModulesByGenericId returns a ModuleMap with all modules with the given
// genericModuleId.
func GetModulesByGenericId(genericModuleId string) ModuleMap {
	return registeredModulesById[genericModuleId]
}

// GetModuleCountByGenericId returns the number of registered modules with the
// given generic id.
func GetModuleCountByGenericId(genericModuleId string) int {
	return len(GetModulesByGenericId(genericModuleId))
}

// GetModuleById returns the Module instance associated with the given
// genericModuleId and SpecificModuleId.
func GetModuleById(genericModuleId, specificModuleId string) Module {
	return registeredModulesById[genericModuleId][specificModuleId]
}

// GetDefaultModuleByGenericId returns the default module represented by the
// given genericModuleId. There may not be a default instance available.
func GetDefaultModuleByGenericId(genericModuleId string) Module {
	// Default module jas the empty string as the specificModuleId.
	return registeredModulesById[genericModuleId][""]
}

// GetAllModules returns a ModuleMap containing all registered modules.
func GetAllModules() FullModuleMap {
	return registeredModulesById
}

// GetAllModulesCount returns the total number of registered modules.
func GetAllModulesCount() int {
	return countFullModuleMap(GetAllModules())
}

// countFullModuleMap returns the number of modules in the given FullModuleMap.
func countFullModuleMap(fullModuleMap FullModuleMap) int {
	count := 0
	for _, s := range fullModuleMap {
		count = count + len(s)
	}

	return count
}
