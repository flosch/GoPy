package vm

import (
	"fmt"
)

type PyObject interface {
	getValue() interface{}
	isTrue() bool
	asString() *string
	
	getattr(name, standard PyObject) PyObject
	operation(op int, obj2 PyObject, inplace bool) (PyObject, PyObject)
	setItem(key, value PyObject) PyObject
	getItem(key PyObject) PyObject
}

type PyObjectData struct {
	attributes map[string]PyObject
}

func (obj *PyObjectData) pyObjInit() {
	obj.attributes = make(map[string]PyObject)
}

func (obj *PyObjectData) operation(op int, obj2 PyObject, inplace bool) (PyObject, PyObject) {
	return PyTypeError, nil
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

// Returns nil or exception object
func (obj *PyObjectData) setItem(key, value PyObject) PyObject {
	if _, ok := obj.attributes["__setitem__"]; ok { // set_func
		// found setattr, call it!
		panic("not implemented yet")	
	}
	panic("stop")
	return PyTypeError
} 

// Returns actual object or exception object
func (obj *PyObjectData) getItem(key PyObject) PyObject {
	if _, ok := obj.attributes["__getitem__"]; ok { // set_func
		// found setattr, call it!
		panic("not implemented yet")	
	}
	return PyTypeError
}
