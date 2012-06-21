package vm

import (
	"fmt"
)

type PyObject interface {
	getValue() interface{}
	isTrue() bool
	asString() *string
	
	getattr(name, standard PyObject) PyObject
}

type PyObjectData struct {
	attributes map[string]PyObject
}

func (obj *PyObjectData) pyObjInit() {
	obj.attributes = make(map[string]PyObject)
}

func (obj *PyObjectData) getattr(name, standard PyObject) PyObject {
	name_string, ok := name.(*PyString)
	if !ok {
		panic(fmt.Sprintf("getattr(_, name [%v]) is no PyString", name))
	}
	value, found := obj.attributes[*name_string.asString()]
	if !found {
		if standard != PyNil {
			return standard
		} else {
			return PyAttributeError
		}
	}
	return value
}