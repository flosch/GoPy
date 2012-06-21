package vm

import (
	"log"
)

const (
	CmpLT = iota
	CmpLTE
	CmpEQ
	CmpNE
	CmpGT
	CmpGTE
	CmpIN
	CmpNIN // not in
	CmpIS
	CmpISN // is not
	CmpEM // Exact match
	CmpBAD //  Bad ???? 
)

var compareMap = map[int]func(obj1, obj2 PyObject) PyObject {
	CmpLT: compareLT,
}

func compareLT(obj1raw, obj2raw PyObject) PyObject {
	obj1, Obj1isInt := obj1raw.(*PyInt)
	obj2, Obj2isInt := obj2raw.(*PyInt)
	
	if Obj1isInt && Obj2isInt {
		if obj1.value < obj2.value {
			return PyTrue
		}
	} else {
		log.Printf("type(obj1) = %T, type(obj2) = %T\n", obj1, obj2)
		panic("Comparison not supported yet")
	}
	
	return PyFalse
}