package vm

type PyArgs struct {
	positional []PyObject
	keyword    map[string]PyObject
}

func PyArgs_Build(args ...PyObject) {
	panic("TODO")
}

func (pa *PyArgs) addPositional(obj PyObject) {
	pa.positional = append(pa.positional, obj)
}

func NewPyArgs() *PyArgs {
	return &PyArgs{
		positional: make([]PyObject, 0, 100),
		keyword:    make(map[string]PyObject),
	}
}
