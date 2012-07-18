package vm

import (
	"fmt"
	"log"
)

type CodeFlags struct {
	optimized,
	newlocals,
	varargs,
	varkeywords,
	nested,
	generator uint32
}

type PyCode struct {
	PyObjectData

	argcount,
	nlocals,
	stacksize,
	raw_flags,
	firstlineno uint32

	consts,
	names,
	varnames,
	freevars,
	cellvars,
	lnotab PyObject

	filename,
	name *string

	flags CodeFlags

	vm   *VM
	code *codeReader
}

func (pc *PyCode) getFromCellFreeStorage(idx int) (obj PyObject) {
	// TODO: Add some range checks

	if idx < len(pc.cellvars.(*PyTuple).items) {
		// variable is cellvars[i] if i is less than the length of cellvars
		obj = pc.cellvars.(*PyTuple).items[idx]
	} else {
		obj = pc.freevars.(*PyTuple).items[idx-len(pc.cellvars.(*PyTuple).items)]
	}
	return
}

func (pc *PyCode) setCellFreeStorage(idx int, item PyObject) error {

	return nil
}

func (pc *PyCode) log(msg string, debug bool) {
	if debug && !pc.vm.debug {
		// Ignore debug messages if debug mode is off
		return
	}
	if debug {
		log.Println(fmt.Sprintf("[%s/%s:DEBUG] %s", *pc.filename, *pc.name, msg))
	} else {
		log.Println(fmt.Sprintf("[%s/%s] %s", *pc.filename, *pc.name, msg))
	}
}

func (pc *PyCode) runtimeError(msg string) {
	panic(fmt.Sprintf("[%s/%s] Runtime error: %s", *pc.filename, *pc.name, msg))
}

func (pc *PyCode) asString() *string {
	return pc.name
}

func (pc *PyCode) getValue() interface{} {
	return pc.code
}

func (pc *PyCode) isTrue() bool {
	return true
}

func (pc *PyCode) getattr(name, standard PyObject) PyObject {
	return nil
}

func NewPyCode(vm *VM) PyObject {
	co := new(PyCode)
	co.pyObjInit()
	co.vm = vm
	co.argcount, _ = vm.content.readDWord()
	co.nlocals, _ = vm.content.readDWord()
	co.stacksize, _ = vm.content.readDWord()
	co.raw_flags, _ = vm.content.readDWord()
	co.code = NewCodeReader([]byte(*vm.readObject().getValue().(*string)))
	co.consts = vm.readObject()
	co.names = vm.readObject()
	co.varnames = vm.readObject()
	co.freevars = vm.readObject()
	co.cellvars = vm.readObject()
	co.filename = vm.readObject().getValue().(*string)
	co.name = vm.readObject().getValue().(*string)
	co.firstlineno, _ = vm.content.readDWord()
	co.lnotab = vm.readObject()
	co.parseFlags()
	return co
}

func (co *PyCode) parseFlags() {
	co.flags.optimized = co.raw_flags & 0x0001
	co.flags.newlocals = co.raw_flags & 0x0002
	co.flags.varargs = co.raw_flags & 0x0004
	co.flags.varkeywords = co.raw_flags & 0x0008
	co.flags.nested = co.raw_flags & 0x0010
	co.flags.generator = co.raw_flags & 0x0020
}
