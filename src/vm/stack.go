package vm

import (
	"log"
)

// PyObjStack
type PyObjStack struct {
	items []PyObject
	count int
}

func NewPyObjStack(count uint64) *PyObjStack {
	return &PyObjStack{
		items: make([]PyObject, count),
	}
}

func (s *PyObjStack) Push(item PyObject) {
	if s.count >= len(s.items) {
		// way better than append() because of continual memory allocation
		if debugMode {
			log.Println("!!! More PyObjStack space required, acquiring...")
		}
		items := make([]PyObject, len(s.items) * 2)
		copy(items, s.items)
		s.items = items
	}
	s.items[s.count] = item
	s.count++
}

func (s *PyObjStack) Pop() PyObject {
	if s.count == 0 {
		return nil
	}
	item := s.items[s.count - 1]
	s.count--
	return item
}

// BlockStack
type Block struct {
	position,
	position2 int64
}

type BlockStack struct {
	items []*Block
	count int
}

func NewBlockStack(count uint64) *BlockStack {
	return &BlockStack{
		items: make([]*Block, count),
	}
}

func (s *BlockStack) Push(position, position2 int64) {
	if s.count >= len(s.items) {
		if debugMode {
			log.Println("!!! More BlockStack space required, acquiring...")
		}
		items := make([]*Block, len(s.items) * 2)
		copy(items, s.items)
		s.items = items
	}
	s.items[s.count] = &Block{
		position: position,
		position2: position2,
	}
	s.count++
}

func (s *BlockStack) Pop() *Block {
	if s.count == 0 {
		return nil
	}
	item := s.items[s.count - 1]
	s.count--
	return item
}