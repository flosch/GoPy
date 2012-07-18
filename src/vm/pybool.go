package vm

type PyBool struct {
	PyObjectData
	value bool
}

func (pb *PyBool) getValue() interface{} {
	return pb.value
}

func (pb *PyBool) isTrue() bool {
	return pb.value == true
}

func (pb *PyBool) asString() *string {
	var str string
	if pb.value {
		str = "True"
	} else {
		str = "False"
	}
	return &str
}

func NewPyBool(value bool) PyObject {
	pb := &PyBool{
		value: value,
	}
	pb.pyObjInit()
	return pb
}
