package vm

import (
	"fmt"
)

type PyList struct {
	PyTuple
}

func (pl *PyList) asString() *string {
	str := fmt.Sprintf("[%s]", *pl.buildItemString())
	return &str
}

// Returns nil or exception object
func (pl *PyList) setItem(key, value PyObject) PyObject {
	if _, ok := pl.attributes["__setitem__"]; ok { // set_func
		// found setattr, call it!
		panic("not implemented yet")
	}
	idxobj, isInt := key.(*PyInt)
	if !isInt {
		fmt.Sprintf("joo\n")
		return PyTypeError
	}

	if int(idxobj.value) > len(pl.items) {
		return PyIndexError
	}

	pl.items[idxobj.value] = value

	return nil
}

// Returns actual object or exception object
func (pl *PyList) getItem(key PyObject) PyObject {
	if _, ok := pl.attributes["__getitem__"]; ok { // set_func
		// found setattr, call it!
		panic("not implemented yet")
	}
	idxobj, isInt := key.(*PyInt)
	if !isInt {
		fmt.Sprintf("joo\n")
		return PyTypeError
	}

	if int(idxobj.value) > len(pl.items) {
		return PyIndexError
	}

	return pl.items[idxobj.value]
}

func NewPyList(items []PyObject) PyObject {
	pl := &PyList{
		PyTuple{items: items},
	}
	pl.pyObjInit()
	return pl
}
