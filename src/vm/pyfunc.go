package vm

import (
	"fmt"
	"log"
	"time"
)

const (
	PyFuncExternal = iota
	PyFuncInternal
)

type PyFunc struct {
	PyObjectData

	// Internal func
	module *Module
	name   *string
	mfunc  ModuleFunc

	// External func
	codeobj PyObject

	// Meta
	functype int
	closure  PyObject // nil or tuple of cell objects
	_doc     PyObject // __doc__ attribute | not used yet
	_name    PyObject // __name__ attribute | not used yet
}

func (pf *PyFunc) setClosure(c PyObject) {
	pf.closure = c
}

func (pf *PyFunc) getValue() interface{} {
	return pf.codeobj
}

func (pf *PyFunc) isTrue() bool {
	return true
}

func (pf *PyFunc) isExternal() bool {
	return pf.functype == PyFuncExternal 
}

func (pf *PyFunc) asString() *string {
	var str string
	switch pf.functype {
	case PyFuncInternal:
		str = fmt.Sprintf("<internal function %s>", *pf.name)
	case PyFuncExternal:
		str = fmt.Sprintf("<external function %s>", *pf.codeobj.asString())
	default:
		panic("unknown func type")
	}
	return &str
}

func (pf *PyFunc) log(msg string) {
	var ident string

	switch pf.functype {
	case PyFuncInternal:
		ident = fmt.Sprintf("%s.%s", pf.module.name, *pf.name)
	case PyFuncExternal:
		ident = fmt.Sprintf("%s/%s", *pf.codeobj.(*PyCode).filename, *pf.codeobj.(*PyCode).name)
	default:
		panic("unknown func type")
	}

	log.Println(fmt.Sprintf("[%s] %s", ident, msg))
}

func (pf *PyFunc) run(args *PyArgs) PyObject {
	// Create frame for run
	frame := NewPyFrame(1000) // TODO change stack size to a better value?

	starttime := time.Now()
	defer func() {
		if debugMode {
			pf.log(fmt.Sprintf("Execution took %s.", time.Since(starttime)))
		}
	}()

	switch pf.functype {
	case PyFuncInternal:
		return pf.mfunc(args)
	case PyFuncExternal:
		if args != nil {
			for i, value := range args.positional {
				// Iterate reverse! Therefore: 
				idx := len(args.positional) - 1 - i

				name := pf.codeobj.(*PyCode).varnames.(*PyTuple).items[idx]
				frame.names[*name.asString()] = value
				fmt.Printf("\n  --- Setting %v -> %v...\n", *name.asString(), *value.asString())
			}
			if len(args.keyword) > 0 {
				panic("Not implemented")
			}
		}
		if debugMode {
			pf.log("Called")
		}
		res, err := pf.codeobj.(*PyCode).eval(frame)
		if err != nil {
			pf.codeobj.(*PyCode).runtimeError(err.Error())
		}
		return res
	default:
		panic("unknown func type")
	}
	panic("unreachable")
}

func NewPyFuncInternal(module *Module, fn ModuleFunc, name *string) PyObject {
	return &PyFunc{
		module:   module,
		name:     name,
		functype: PyFuncInternal,
		mfunc:    fn,
	}
}

func NewPyFuncExternal(codeobj PyObject) PyObject {
	pf := &PyFunc{
		codeobj:  codeobj,
		functype: PyFuncExternal,
	}
	pf.pyObjInit()
	return pf
}
