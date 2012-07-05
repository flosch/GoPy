package vm

import (
	//"log"
)

var Modules = map[string]Module{
	"time": ModuleTime,
	"gopy": ModuleGopy,
}

type ModuleFunc func(args *PyArgs) PyObject
type ModuleDict map[string]ModuleFunc

type Module struct {
	name string
	funcs ModuleDict
}

func (m *Module) inject(pm *PyModule) error {
	for name, fn := range m.funcs {
		//log.Printf("Injecting %v -> %v\n", name, fn)
		pm.attributes[name] = NewPyFuncInternal(m, fn, &name)
	}
	return nil
}