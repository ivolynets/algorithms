package collection

// Stack initial capacity
const initCap = 8

// Array stack type
type ArrayStack struct {
	values []interface{}
	size   int
}

// Creates a new array stack instance and returns a pointer to it.
func NewArrayStack() *ArrayStack {
	return &ArrayStack{make([]interface{}, initCap), 0}
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

	if l := len(s.values); l > initCap {

		c := s.size + (s.size >> 1)
		if c < initCap {
			c = initCap
		}

		if l-c >= l>>2 {

			d := make([]interface{}, c)
			copy(d, s.values[:s.size])

			s.values = d

		}

	}

}
