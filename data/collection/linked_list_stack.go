package collection

// Internal stack item representation
type stackItem struct {
	value interface{}
	next  *stackItem
}

// Linked list stack type
type LinkedListStack struct {
	head *stackItem
	size int
}

// Creates a new linked list stack and returns a pointer to it.
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil, 0}
}

// Adds a new item at the top of the stack.
func (s *LinkedListStack) Push(i interface{}) {

	si := &stackItem{i, s.head}
	s.head = si
	s.size++

}

// Returns but does not remove an item at the top of the stack. If stack is not
// empty than i contains value and ok is true. Otherwise i contains nil and ok
// is false.
func (s *LinkedListStack) Peek() (i interface{}, ok bool) {

	if ok = s.size > 0; ok {
		i = s.head.value
	} else {
		i = nil
	}

	return

}

// Removes and returns an item at the top of the stack. If stack is not empty
// than i contains value and ok is true. Otherwise i contains nil and ok is
// false.
func (s *LinkedListStack) Pop() (i interface{}, ok bool) {

	if s.size > 0 {

		i = s.head.value
		s.head = s.head.next
		s.size--

		ok = true

	} else {

		i = nil
		ok = false

	}

	return

}

// Returns size of the stack.
func (s *LinkedListStack) Size() int {
	return s.size
}
