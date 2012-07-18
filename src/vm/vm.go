package vm

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type Runtime struct {
	starttime, endtime time.Time

	mainframe    *PyFrame
	freevars     []PyObject // TODO: better place?! length!? is usage right?!
	instructions uint64

	running bool
	lock    sync.RWMutex
}

type VM struct {
	filename         string
	debug            bool
	content          *codeReader
	code             PyObject
	interned_strings []PyObject

	runtime Runtime
}

func (vm *VM) log(msg string) {
	log.Println(fmt.Sprintf("[VM] %s", msg))
}

func (vm *VM) parse() error {
	log.Println("Parsing...")
	if magic, _ := vm.content.readDWord(); magic != 168686339 {
		log.Fatal("No valid compiled python file (invalid magic)")
	}

	timestamp, _ := vm.content.readDWord()
	t := time.Unix(int64(timestamp), 0)
	log.Printf("File created: %s (timestamp: %d)\n", t, timestamp)

	vm.interned_strings = make([]PyObject, 0, 5000) // TODO: Wahllose Kapazität besser bestimmen!
	vm.code = vm.readObject()

	log.Println("Parsing finished")
	return nil
}

func (vm *VM) Filename() *string {
	return vm.code.(*PyCode).filename
}

func (vm *VM) Name() *string {
	return vm.code.(*PyCode).name
}

func (vm *VM) Run() error {
	log.Println("Running...")

	vm.log(fmt.Sprintf("Stacksize = %d", vm.code.(*PyCode).stacksize))
	vm.runtime.mainframe = NewPyFrame(uint64(vm.code.(*PyCode).stacksize))
	vm.runtime.starttime = time.Now()

	if retval, err := vm.code.(*PyCode).eval(vm.runtime.mainframe); err != nil {
		return err
	} else {
		vm.log(fmt.Sprintf("Returning value: %v (%T)", *retval.asString(), retval))
	}

	vm.runtime.endtime = time.Now()
	vm.log(fmt.Sprintf("Execution of program took %s.", vm.runtime.endtime.Sub(vm.runtime.starttime)))

	log.Printf("Running finished (%d instructions ran).\n", vm.runtime.instructions)
	return nil
}

var debugMode bool = false

func NewVM(filename string, debug bool) (*VM, error) {
	debugMode = debug

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	vm := &VM{
		content:  NewCodeReader(content),
		filename: filename,
		debug:    debug,
		runtime: Runtime{
			freevars: make([]PyObject, 1000, 1000),
		},
	}

	if err := vm.parse(); err != nil {
		return nil, err
	}

	return vm, nil
}
