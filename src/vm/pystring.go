package vm

type PyString struct {
	PyObjectData
	value *string
}

func (ps *PyString) getValue() interface{} {
	return ps.value
}

func (ps *PyString) isTrue() bool {
	return len(*ps.value) > 0
}

func (ps *PyString) asString() *string {
	return ps.value
}

func NewPyString(value *string) PyObject {
	ps := &PyString{
		value: value,
	}
	ps.pyObjInit()
	return ps
}