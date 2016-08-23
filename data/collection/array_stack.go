package collection

// Stack initial capacity
const initStackCap = 10

// Array stack type
type ArrayStack struct {
	values []interface{}
	size   int
}

// Creates a new array stack instance and returns a pointer to it.
func NewArrayStack() *ArrayStack {
	return &ArrayStack{make([]interface{}, initStackCap), 0}
}

// Adds a new item at the top of the stack.
func (s *ArrayStack) Push(i interface{}) {

	s.extendStack(s.size + 1)
	s.values[s.size] = i
	s.size++

}

// Returns but does not remove an item at the top of the stack. If stack is not
// empty than i contains value and ok is true. Otherwise i contains nil and ok
// is false.
func (s *ArrayStack) Peek() (i interface{}, ok bool) {

	if ok = s.size > 0; ok {
		i = s.values[s.size-1]
	} else {
		i = nil
	}

	return

}

// Removes and returns an item at the top of the stack. If stack is not empty
// than i contains value and ok is true. Otherwise i contains nil and ok is
// false.
func (s *ArrayStack) Pop() (i interface{}, ok bool) {

	if s.size > 0 {

		s.size--
		i = s.values[s.size]

		s.compactStack()
		ok = true

	} else {

		i = nil
		ok = false

	}

	return

}

// Returns size of the stack.
func (s *ArrayStack) Size() int {
	return s.size
}

// Increases stack capacity if it's not enough.
func (s *ArrayStack) extendStack(c int) {

	if len(s.values) < c {

		d := make([]interface{}, s.size+(s.size>>1))
		copy(d[:s.size], s.values)

		s.values = d

	}

}

// Decreases stack capacity if it's too large.
func (s *ArrayStack) compactStack() {

	if l := len(s.values); l > initStackCap {

		c := l - l/3
		if c < initStackCap {
			c = initStackCap
		}

		if c >= s.size {

			d := make([]interface{}, c)
			copy(d, s.values[:s.size])

			s.values = d

		}

	}

}
