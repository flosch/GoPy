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

func NewPyList(items []PyObject) PyObject {
	pl := &PyList{
		PyTuple{items: items},
	}
	pl.pyObjInit()
	return pl
}