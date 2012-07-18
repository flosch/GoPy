package vm

var PyAttributeError = NewPyException("AttributeError")
var PyNameError = NewPyException("NameError")
var PyTypeError = NewPyException("TypeError")
var PyIndexError = NewPyException("IndexError")

type PyException struct {
	PyObjectData
	name *string
	msg  *string
}

func (pe *PyException) asString() *string {
	return pe.name
}

func (pe *PyException) getValue() interface{} {
	return pe
}

func (pe *PyException) isTrue() bool {
	return true
}

func NewPyException(name string) PyObject {
	excp := new(PyException)
	excp.pyObjInit()
	excp.name = &name
	return excp
}
