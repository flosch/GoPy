package vm

type PyFrame struct {
	stack *PyObjStack
	blocks *BlockStack
	names map[string]PyObject
	position int64
	funcs map[string]PyFunc
}

func NewPyFrame(stacksize uint64) *PyFrame {
	return &PyFrame{
		stack: NewPyObjStack(stacksize),
		blocks: NewBlockStack(10000),
		names: make(map[string]PyObject),
	}
}