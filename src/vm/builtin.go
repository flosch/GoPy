package vm

var PyTrue = NewPyBool(true)
var PyFalse = NewPyBool(false)
var PyNil = NewPyNone()

var PyBuiltInTypeMap = map[string]PyObject{
	"True": PyTrue,
	"False": PyFalse,
	//"None": PyNil,
}

var PyBuiltInFuncMap = map[string]func(){

} 