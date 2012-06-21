package vm

type PyNone struct {
	PyObjectData
}

func (pn *PyNone) getValue() interface{} {
	return nil
}

func (pn *PyNone) isTrue() bool {
	return false
}

func (pn *PyNone) asString() *string {
	str := "None"
	return &str
}

func NewPyNone() PyObject {
	pn := new(PyNone)
	pn.pyObjInit()
	return pn
}