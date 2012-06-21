package vm

import (
	"fmt"
)

type PyModule struct {
	PyObjectData
	module *Module
	name *string
	
	// If module is extern (it's own pyc file):
	content *codeReader	
	code PyObject
	interned_strings []PyObject
}

func (pm *PyModule) asString() *string {
	str := fmt.Sprintf("<module %s>", *pm.name)
	return &str
}

func (pm *PyModule) getValue() interface{} {
	return nil
}

func (pm *PyModule) isTrue() bool {
	return true
}

func NewPyModule(name *string) PyObject {
	mod := new(PyModule)
	mod.pyObjInit()
	mod.name = name
	
	// Import all functions and global names and make them
	// available in the attributes
	module, is_builtin := Modules[*name]
	if is_builtin {
		mod.module = &module 
		if err := module.inject(mod); err != nil {
			panic("Error during module injection: " + err.Error())
		}
	} else {
		// Search for a pyc file and execute it!
		panic(fmt.Sprintf("Non-builtin modules are not supported yet (%v)", *name))
	}
	
	return mod
}