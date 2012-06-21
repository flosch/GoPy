package vm

import (
	//"errors"
	"time"
	"fmt"
	"io"
	//"reflect"
)

const (
	POP_TOP = 1
	ROT_TWO = 2
	BINARY_MULTIPLY = 20
	BINARY_ADD = 23
	INPLACE_ADD = 55
	PRINT_ITEM = 71
	PRINT_NEWLINE = 72
	RETURN_VALUE = 83
	POP_BLOCK = 87
	STORE_NAME = 90
	DELETE_NAME = 91
	UNPACK_SEQUENCE = 92
	FOR_ITER = 93
	LOAD_CONST = 100
	LOAD_NAME = 101
	BUILD_TUPLE = 102
	BUILD_LIST = 103
	LOAD_ATTR = 106
	COMPARE_OP = 107
	IMPORT_NAME = 108
	JUMP_ABSOLUTE = 113
	POP_JUMP_IF_FALSE = 114
	LOAD_GLOBAL = 116
	SETUP_LOOP = 120
	LOAD_FAST = 124
	STORE_FAST = 125
	CALL_FUNCTION = 131
	MAKE_FUNCTION = 132
	MAKE_CLOSURE = 134
	LOAD_CLOSURE = 135
	LOAD_DEREF = 136
	STORE_DEREF = 137
)

const HasArgLimes = 90

const (
	OpAdd = iota
	OpMultiply
)

func (code *PyCode) eval(frame *PyFrame) (PyObject, error) {
	starttime := time.Now()
	defer func() {
		code.log(fmt.Sprintf("Evaluation took %s.", time.Since(starttime)), true) 
	}()

	for {
		// Get opcode
		op_position := frame.position
		code.code.setPos(frame.position)
		opcode, err := code.code.readByte()
		if err != nil {
			if err == io.EOF {
				code.log("End of code reached", false)
				break
			} else {
				panic(err.Error())
			}
		}
		frame.position += 1
		code.vm.runtime.instructions += 1
				
		// Has opcode arguments?
		var oparg uint16
		if opcode >= HasArgLimes {
			oparg, err = code.code.readWord()
			if err != nil {
				panic("opcode requires arguments, but fetching failed: " + err.Error())
			}
			frame.position += 2
		}
		
		switch opcode {
			case POP_TOP:
				code.log("Pop top", true)
				frame.stack.Pop()
			case ROT_TWO:
				code.log("Rot two", true)
				op1stack := frame.stack.Pop()
				if op1stack == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				op2stack := frame.stack.Pop()
				if op2stack == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				frame.stack.Push(op1stack)
				frame.stack.Push(op2stack)
			case BINARY_MULTIPLY:
				code.log("Binary add/inplace add", true)
				op1 := frame.stack.Pop()
				if op1 == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				
				op2 := frame.stack.Pop()
				if op2 == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}

				result := op2.operation(OpMultiply, op1)
				if _, isException := result.(*PyException); isException {
					code.runtimeError("Exception raised!") // TODO: handle correctly, this is only provisory
				}
				code.log(fmt.Sprintf("Multiplying: %s * %s = %s", *op2.asString(), *op1.asString(), *result.asString()), true) 
				frame.stack.Push(result)
			case INPLACE_ADD, BINARY_ADD:
				code.log("Binary add/inplace add", true)
				op1stack := frame.stack.Pop()
				if op1stack == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				op1, ok := op1stack.(*PyInt)
				
				if !ok {
					code.runtimeError("Object must be an PyInt")
				}
				
				op2stack := frame.stack.Pop()
				if op2stack == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				op2, ok := op2stack.(*PyInt)
				
				if !ok {
					code.runtimeError("Object must be an PyInt")
				}
				
				result := NewPyInt(op2.value + op1.value)
				code.log(fmt.Sprintf("Adding: %d + %d = %v", op2.value, op1.value, result.getValue()), true) 
				frame.stack.Push(result)
			case PRINT_ITEM:
				code.log("Print item", true)
				stackitem := frame.stack.Pop()
				if stackitem == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				//fmt.Printf("Item: %v\n", stackitem)
				fmt.Printf("%s ", *stackitem.asString())
			case PRINT_NEWLINE:
				fmt.Println()
				code.log("Print newline", true)
			case UNPACK_SEQUENCE:
				code.log("Unpack sequence", true)
				stackitem := frame.stack.Pop()
				if stackitem == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				
				items := stackitem.getValue().([]PyObject)
				for i := len(items) - 1; i >= 0; i-- {
					code.log(fmt.Sprintf("Unpacking %d -> %v", i, items[i].getValue()), true)
					frame.stack.Push(items[i])
				}
			case RETURN_VALUE:
				// TODO: Abschlussarbeiten durchführen, was zB? Siehe ceval.c in CPython
				code.log("Return value", true)
				stackitem := frame.stack.Pop()
				if stackitem == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				return stackitem, nil
			case POP_BLOCK:
				code.log("Pop block", true)
				frame.blocks.Pop()
			case LOAD_NAME:
				name := *code.names.(*PyTuple).getItem(int(oparg)).getValue().(*string)
				
				// Check wheter it's a built in type
				value, is_builtin := PyBuiltInTypeMap[name]
				if is_builtin {
					code.log(fmt.Sprintf("Load built-in name %s (= %v, push to stack)", name, *value.asString()), true)
					frame.stack.Push(value)
				} else {
					// It's no built-in, determine value
					
					_, global_found := code.vm.runtime.mainframe.names[name]
					
					// TODO CHECK: Not sure whether to give global or local priority 
					if global_found {
						// Get global 
					} else {
					
					}
					
					// Workaround: Get from local context
					value, local_found := frame.names[name]
					if !local_found {
						code.runtimeError(fmt.Sprintf("Could not find name (%v) in local namespace", name))
					}
					code.log(fmt.Sprintf("Load name %s (= %v, push to stack)", name, *value.asString()), true)
					frame.stack.Push(value)
				}
			case STORE_NAME:
				name := *code.names.(*PyTuple).getItem(int(oparg)).getValue().(*string)
				
				_, global_found := code.vm.runtime.mainframe.names[name]
				
				// TODO CHECK: Not sure whether to give global or local priority 
				if global_found {
					// Set global 
				} else {
				
				}
				
				// Workaround: Set in local context
				stackitem := frame.stack.Pop()
				if stackitem == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				frame.names[name] = stackitem
				
				code.log(fmt.Sprintf("Store name: %s = %v", name, *frame.names[name].asString()), true)
			case STORE_FAST:
				stackitem := frame.stack.Pop()
				if stackitem == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				name := *code.varnames.(*PyTuple).getItem(int(oparg)).(*PyString).asString()
				frame.names[name] = stackitem
				code.log(fmt.Sprintf("Store FAST name: %s = %v", name, *frame.names[name].asString()), true)
			case LOAD_FAST:
				name := *code.varnames.(*PyTuple).getItem(int(oparg)).(*PyString).asString()
				item, ok := frame.names[name]
				if !ok {
					code.runtimeError("Could not find item in varnames")
				}
				code.log(fmt.Sprintf("Load FAST name: %s (= %v, pushing on stack)", name, *frame.names[name].asString()), true)
				frame.stack.Push(item)
			case BUILD_TUPLE:
				items := make([]PyObject, oparg)
				for i := 0; i < int(oparg); i++ {
					stackitem := frame.stack.Pop()
					if stackitem == nil {
						code.runtimeError("Stackitem is nil during tuple build process")
					}
					items[i] = stackitem
				}
				tuple := NewPyTuple(items)
				frame.stack.Push(tuple)
				code.log(fmt.Sprintf("Build tuple (%d items: %s)", oparg, *tuple.(*PyTuple).asString()), true)
			case BUILD_LIST:
				items := make([]PyObject, oparg)
				for i := 0; i < int(oparg); i++ {
					stackitem := frame.stack.Pop()
					if stackitem == nil {
						code.runtimeError("Stackitem is nil during list build process")
					}
					items[i] = stackitem 
				}
				list := NewPyList(items)
				frame.stack.Push(list)
				code.log(fmt.Sprintf("Build list (%d items: %s)", oparg, *list.(*PyList).asString()), true)
			case LOAD_ATTR:
				name := code.names.(*PyTuple).getItem(int(oparg)).(*PyString)
				obj := frame.stack.Pop()
				result := obj.getattr(name, PyNil) // TODO: Check+raise Exception in result! This might return PyAttributeError
				
				if _, is_exception := result.(*PyException); is_exception {
					panic("Exception raised! To be handled correctly.")
				}

				frame.stack.Push(result)
				code.log(fmt.Sprintf("Load attr [getattr(%v, %v) = %v]", *obj.asString(), *name.asString(), *result.asString()), true)
			case COMPARE_OP:
				op1stack := frame.stack.Pop()
				if op1stack == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				op2stack := frame.stack.Pop()
				if op2stack == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				compareFunc, found := compareMap[int(oparg)]
				if !found {
					code.runtimeError(fmt.Sprintf("Could not find compare function %d in map", oparg))
				}
				result := compareFunc(op2stack, op1stack)
				code.log(fmt.Sprintf("Compare op (%d); result = %t", oparg, result.isTrue()), true)
				frame.stack.Push(result)
			case LOAD_CONST:
				value := code.consts.(*PyTuple).getItem(int(oparg))
				code.log(fmt.Sprintf("Load const: %v (pushing on stack)", *value.asString()), true)
				frame.stack.Push(value)
			case LOAD_GLOBAL:
				name := *code.names.(*PyTuple).getItem(int(oparg)).getValue().(*string)
				
				// Check wheter it's a built in type
				value, is_builtin := PyBuiltInTypeMap[name]
				if is_builtin {
					code.log(fmt.Sprintf("Load built-in name %s (= %v, push to stack)", name, *value.asString()), true)
					frame.stack.Push(value)
				} else {
					// It's no built-in, determine value
					value, global_found := code.vm.runtime.mainframe.names[name]
					
					// TODO CHECK: Not sure whether to give global or local priority 
					if !global_found {
						code.runtimeError(fmt.Sprintf("Could not find name (%v) in global namespace", name))
					}
					
					code.log(fmt.Sprintf("Load GLOBAL name %s (= %v, push to stack)", name, *value.asString()), true)
					frame.stack.Push(value)
				}
			case SETUP_LOOP:
				code.log("Setup loop", true)
				frame.blocks.Push(op_position, frame.position + int64(oparg))
			case IMPORT_NAME:
				name := code.names.(*PyTuple).getItem(int(oparg)).(*PyString).asString()
				module := NewPyModule(name)
				frame.stack.Push(module)
				code.log(fmt.Sprintf("Import name (%s = %v, pushed on stack)", *name, *module.asString()), true)
			case JUMP_ABSOLUTE:
				code.log(fmt.Sprintf("Jump absolute (%d)\n", oparg), true)
				frame.position = int64(oparg)
			case POP_JUMP_IF_FALSE:
				stackitem := frame.stack.Pop()
				if stackitem == nil {
					code.runtimeError("Stackitem cannot be nil.")
				}
				
				if !stackitem.isTrue() {
					frame.position = int64(oparg)
				} 
				
				code.log(fmt.Sprintf("Pop jump if false (result = %t)", stackitem.isTrue()), true)
			case CALL_FUNCTION:
				code.log(fmt.Sprintf("Call function (args=%d)", oparg), true)
				if oparg > 0 {
					nkwargs := (oparg >> 8) & 0xff
					// Keyword arguments first (high byte)
					for i := 0; i < int(nkwargs); i++ {
						arg := frame.stack.Pop()
						code.log(fmt.Sprintf("   Received kw-arg: %v", arg), true)
					}			
							
					nposargs := oparg & 0xff 
					// positional parameters (low byte)
					for i := 0; i < int(nposargs); i++ {
						arg := frame.stack.Pop()
						code.log(fmt.Sprintf("   Received pos-arg: %v", arg), true)
					}
				}
				fobj := frame.stack.Pop()
				if fobj == nil {
					code.runtimeError("Stack empty, expected: function object")
				}
				result := fobj.(*PyFunc).run(nil) // TODO FIX!!
				frame.stack.Push(result) // dunno?
			case MAKE_FUNCTION:
				code.log(fmt.Sprintf("Make function (argcount=%d)", oparg), true)
				codeobj := frame.stack.Pop()
				fobj := NewPyFuncExternal(codeobj)
				if oparg > 0 {
					// Parse arguments
					panic("Not implemented yet")
					for i := 0; i < int(oparg); i++ {
						_ = frame.stack.Pop()
					}
				}
				frame.stack.Push(fobj)
			case MAKE_CLOSURE:
				code.log(fmt.Sprintf("Make closure (argc = %d)", oparg), true)
				c := frame.stack.Pop()
				if c == nil {
					code.runtimeError("No code for closure (stackitem is nil)")
				}
				fn := NewPyFuncExternal(code)
				closure := frame.stack.Pop()
				if closure == nil {
					code.runtimeError("No closure for make closure (stackitem is nil)")
				}
				fn.(*PyFunc).setClosure(closure)
				
				if oparg > 0 {
					// untested, but should work
					items := make([]PyObject, oparg)
					for i := 0; i < int(oparg); i++ {
						stackitem := frame.stack.Pop()
						if stackitem == nil {
							code.runtimeError("Stackitem is nil during tuple build process in make closure")
						}
						items = append(items, stackitem) 
					}
					tuple := NewPyTuple(items)
					frame.stack.Push(tuple)
				}
				
				frame.stack.Push(fn)
			case LOAD_CLOSURE:
				panic("Check for correct implementation")
				obj := code.getFromCellFreeStorage(int(oparg))
				frame.stack.Push(obj)
				code.log(fmt.Sprintf("Load closure (%s)", *obj.(*PyString).asString()), true)
			case LOAD_DEREF:
				panic("Check for correct implementation")
				code.log(fmt.Sprintf("Load deref (idx = %d)", oparg), true)
				obj := code.vm.runtime.freevars[int(oparg)]
				if obj == nil {
					code.runtimeError("Received freevar was nil")
				}
				frame.stack.Push(obj)
			case STORE_DEREF:
				panic("Check for correct implementation")
				code.log(fmt.Sprintf("Store deref (idx = %d)", oparg), true)
				obj := frame.stack.Pop()
				code.vm.runtime.freevars[int(oparg)] = obj 
			default:
				code.runtimeError(fmt.Sprintf("!!! Unhandled opcode: %c (%d)", opcode, opcode))
		}
	}
	//return PyNil, nil
	panic("unreachable")
}