package vm

import (
	"fmt"
	"log"
)

func (vm *VM) readObject() PyObject {
	c, err := vm.content.readByte()
	if err != nil {
		log.Println(err.Error())
	}
	if vm.debug {
		log.Printf("Token = '%c' (%d)\n", c, c)
	}
	switch c {
	case 'c': // CodeObject
		return NewPyCode(vm)
	case 's', 't': // String
		size, err := vm.content.readDWord()
		if err != nil {
			log.Fatal(err.Error())
		}

		str, err := vm.content.readString(size)
		if err != nil {
			log.Fatal(err.Error())
		}

		pystr := NewPyString(str)

		if c == 't' {
			// Handle interned string
			vm.interned_strings = append(vm.interned_strings, pystr)
		}

		return pystr
	case '(': // Tuple
		itemcount, err := vm.content.readDWord()
		if err != nil {
			log.Fatal(err.Error())
		}
		var tuple_items []PyObject = make([]PyObject, itemcount)
		for i := 0; i < int(itemcount); i++ {
			tuple_items[i] = vm.readObject()
		}
		return NewPyTuple(tuple_items)
	case 'i': // Integer
		value, err := vm.content.readDWord()
		if err != nil {
			log.Fatal(err.Error())
		}
		return NewPyInt(int64(value))
	case 'N': // None
		return NewPyNone()
	case 'R': // StringRef
		n, err := vm.content.readDWord()
		if err != nil {
			log.Fatal(err.Error())
		}
		//log.Println("Referenced interned string: " + *vm.interned_strings[n].asString())
		return vm.interned_strings[n]
	default:
		log.Fatal(fmt.Sprintf("Unhandled token: '%c' (%d)\n", c, c))
		return nil
	}

	panic("unreachable")
}
