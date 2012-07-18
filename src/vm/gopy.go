package vm

var ModuleGopy = Module{
	funcs: ModuleDict{
		"go": func(args *PyArgs) PyObject {
			return PyTrue
		},
	},
}
