package vm

import (
	"strconv"
)

type PyInt struct {
	PyObjectData
	value int64
}

func (pi *PyInt) getValue() interface{} {
	return pi.value
}

func (pi *PyInt) isTrue() bool {
	return pi.value == 1
}

func (pi *PyInt) asString() *string {
	s := strconv.Itoa(int(pi.value))
	return &s
}

func NewPyInt(value int64) PyObject {
	pi := &PyInt{
		value: value,
	}
	pi.pyObjInit()
	return pi
}