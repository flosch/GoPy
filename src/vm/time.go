package vm

import (
	"time"
)

var ModuleTime = Module{
	name: "time",
	funcs: ModuleDict{
		"sleep": func(args *PyArgs) PyObject {
			time.Sleep(1 * time.Second)
			return PyTrue
		},
		"time": func(args *PyArgs) PyObject {
			return NewPyInt(time.Now().Unix())
		},
	},
}
