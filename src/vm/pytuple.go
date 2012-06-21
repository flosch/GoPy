package vm

import (
	"fmt"
	"strings"
)

type PyTuple struct {
	PyObjectData
	items []PyObject
}

func (pt *PyTuple) getValue() interface{} {
	return pt.items
}

func (pt *PyTuple) getItem(idx int) PyObject {
	return pt.items[idx]
}

func (pt *PyTuple) length() int {
	return len(pt.items)
}

func (pt *PyTuple) isTrue() bool {
	return len(pt.items) > 0
}

func (pt *PyTuple) asString() *string {
	// TODO: Make it more performant
	str := fmt.Sprintf("(%s)", strings.Join(func() []string {
		res := make([]string, len(pt.items))
		for idx, item := range pt.items {
			res[idx] = *item.asString()	 
		}
		return res
	}(), ", "))
	return &str
}

func NewPyTuple(items []PyObject) PyObject {
	pt := &PyTuple{
		items: items,
	}
	pt.pyObjInit()
	return pt
}