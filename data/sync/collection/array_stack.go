package collection

import (
	"sync"

	"github.com/ivolynets/algorithms/data/collection"
)

// Synchronized array stack type
type ArrayStack struct {
	stack *collection.ArrayStack
	mux   sync.RWMutex
}

// Creates a new synchronized array stack instance and returns a pointer to it.
func NewArrayStack() *ArrayStack {
	s := collection.NewArrayStack()
	return &ArrayStack{s, sync.RWMutex{}}
}

// Adds a new item at the top of the stack.
func (s *ArrayStack) Push(i interface{}) {

	s.mux.Lock()
	defer s.mux.Unlock()

	s.stack.Push(i)

}

// Returns but does not remove an item at the top of the stack. If stack is not
// empty than i contains value and ok is true. Otherwise i contains nil and ok
// is false.
func (s *ArrayStack) Peek() (i interface{}, ok bool) {

	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.stack.Peek()

}

// Removes and returns an item at the top of the stack. If stack is not empty
// than i contains value and ok is true. Otherwise i contains nil and ok is
// false.
func (s *ArrayStack) Pop() (i interface{}, ok bool) {

	s.mux.Lock()
	defer s.mux.Unlock()

	return s.stack.Pop()

}

// Returns size of the stack.
func (s *ArrayStack) Size() int {

	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.stack.Size()

}
