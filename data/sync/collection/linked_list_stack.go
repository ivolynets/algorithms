package collection

import (
	"sync"

	"github.com/ivolynets/algorithms/data/collection"
)

// Synchronized linked list stack type
type LinkedListStack struct {
	stack *collection.LinkedListStack
	mux   sync.RWMutex
}

// Creates a new synchronized linked list stack instance and returns a pointer
// to it.
func NewLinkedListStack() *LinkedListStack {
	s := collection.NewLinkedListStack()
	return &LinkedListStack{s, sync.RWMutex{}}
}

// Adds a new item at the top of the stack.
func (s *LinkedListStack) Push(i interface{}) {

	s.mux.Lock()
	defer s.mux.Unlock()

	s.stack.Push(i)

}

// Returns but does not remove an item at the top of the stack. If stack is not
// empty than i contains value and ok is true. Otherwise i contains nil and ok
// is false.
func (s *LinkedListStack) Peek() (i interface{}, ok bool) {

	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.stack.Peek()

}

// Removes and returns an item at the top of the stack. If stack is not empty
// than i contains value and ok is true. Otherwise i contains nil and ok is
// false.
func (s *LinkedListStack) Pop() (i interface{}, ok bool) {

	s.mux.Lock()
	defer s.mux.Unlock()

	return s.stack.Pop()

}

// Returns size of the stack.
func (s *LinkedListStack) Size() int {

	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.stack.Size()

}
