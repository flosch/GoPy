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

func (pt *PyTuple) isTrue() bool {
	return len(pt.items) > 0
}

func (pt *PyTuple) asString() *string {
	str := fmt.Sprintf("(%s)", *pt.buildItemString())
	return &str
}

func (pt *PyTuple) buildItemString() *string {
	// TODO: Make it more performant
	str := strings.Join(func() []string {
		res := make([]string, MinInt(len(pt.items), 11))
		for idx, item := range pt.items {
			res[idx] = *item.asString()
			if idx >= 9 && (idx+1) < (len(pt.items) - 1) {
				res[idx+1] = fmt.Sprintf("... %d more", len(pt.items) - 10)
				break
			}
		}
		return res
	}(), ", ")
	return &str
}

// Returns (PyException, resultobj)
func (pt *PyTuple) operation(op int, obj2 PyObject, inplace bool) (PyObject, PyObject) {
	if !inplace {
		panic("Not implemented")
	}
	
	switch op {
		case OpMultiply:
			value, isInt := obj2.(*PyInt)
			if !isInt {
				fmt.Println("TypeError! Multiply on list/tuple can only be done with integers.")
				return PyTypeError, nil
			}
			
			newList := make([]PyObject, 0, len(pt.items) * int(value.value))
			for i := 0; i < int(value.value); i++ {
				//fmt.Printf("i = %d\n", i)
				for _, item := range pt.items {
					//fmt.Printf("%d = %v\n", i*(idx+1), *item.asString())
					newList = append(newList, item) 
				}
			}
			pt.items = newList
			
			//fmt.Printf("New size: %d, requested = %d\n", pt.length(), int(value.value) * oldSize)
			
			return nil, pt
	}
	return PyTypeError, nil
} 

func NewPyTuple(items []PyObject) PyObject {
	pt := &PyTuple{
		items: items,
	}
	pt.pyObjInit()
	return pt
}