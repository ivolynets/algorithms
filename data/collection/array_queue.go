package collection

// Queue initial capacity
const initQueueCap = 10

// Array queue type
type ArrayQueue struct {
	values []interface{}
	head   int
	tail   int
}

// Creates a new array queue instance and returns a pointer to it.
func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{make([]interface{}, initQueueCap), 0, 0}
}

// Adds a new item at the tail of the queue.
func (s *ArrayQueue) Add(i interface{}) {

	s.extendQueue(s.tail + 1)
	s.values[s.tail] = i
	s.tail++

}

// Returns but does not remove an item at the head of the queue. If queue is not
// empty than i contains value and ok is true. Otherwise i contains nil and ok
// is false.
func (s *ArrayQueue) Peek() (i interface{}, ok bool) {

	if ok = s.tail > s.head; ok {
		i = s.values[s.head]
	} else {
		i = nil
	}

	return

}

// Removes and returns an item at the head of the queue. If queue is not empty
// than i contains value and ok is true. Otherwise i contains nil and ok is
// false.
func (s *ArrayQueue) Poll() (i interface{}, ok bool) {

	if s.tail > s.head {

		i = s.values[s.head]
		s.head++

		s.compactQueue()
		ok = true

	} else {

		i = nil
		ok = false

	}

	return

}

// Returns size of the queue.
func (s *ArrayQueue) Size() int {
	return s.tail - s.head
}

// Increases queue capacity if it's not enough.
func (s *ArrayQueue) extendQueue(c int) {

	l := len(s.values)

	if l < c {

		// if we really need more capacity then extend underlying array

		d := make([]interface{}, l+(l>>1))
		copy(d[:s.Size()], s.values[s.head:s.tail])

		s.values = d

	} else if s.tail >= l {

		// if capacity is enough but we reached the end of the underlying array
		// then rearrange it

		t := s.Size()
		copy(s.values[:t], s.values[s.head:s.tail])
		s.head = 0
		s.tail = t

	}

}

// Decreases queue capacity if it's too large.
func (s *ArrayQueue) compactQueue() {

	if l := len(s.values); l > initQueueCap {

		c := l - l/3
		if c < initQueueCap {
			c = initQueueCap
		}

		if c >= s.Size() {

			d := make([]interface{}, c)
			copy(d, s.values[s.head:s.tail])

			s.values = d

		}

	}

}
