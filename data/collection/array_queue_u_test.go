package collection

import "testing"

// Test the queue internal extension function.
func TestArrayQueueExtend(t *testing.T) {

	// init queue

	lenExp := 3
	capExp := 4

	s := &ArrayQueue{make([]interface{}, capExp), 0, 0}

	for i := 1; i <= lenExp; i++ {
		s.values[i] = i
	}

	s.head = 1
	s.tail = capExp

	// test if initial capacity of underlying array was not changed

	testArrayQueueAttributes(t, "Initial queue", s, 1, capExp, capExp)

	// test how arrangement works

	s.extendQueue(lenExp + 1)
	testArrayQueueAttributes(t, "Arranged queue", s, 0, lenExp, capExp)

	// test how extension works

	capExp = 6
	s.extendQueue(lenExp + 2)
	testArrayQueueAttributes(t, "Extended queue", s, 0, lenExp, capExp)

	// test if all items in the queue are still in place

	for i := 0; i < s.Size(); i++ {
		if s.values[i] != i+1 {
			t.Fatalf("Extended queue: item %v not found", i+1)
		}
	}

}

// Test the queue internal compaction function.
func TestArrayQueueCompact(t *testing.T) {

	// init queue

	lenExp := initQueueCap + 1
	capExp := initQueueCap + initQueueCap>>1
	capExp = capExp + capExp>>1

	s := &ArrayQueue{make([]interface{}, capExp), 0, 0}

	for i := 0; i < lenExp; i++ {
		s.values[i] = i + 1
		s.tail++
	}

	// test if initial capacity of underlying array was not changed

	testArrayQueueAttributes(t, "Initial queue", s, 0, lenExp, capExp)

	// test how copaction works

	capExp = initQueueCap + initQueueCap>>1
	s.compactQueue()
	testArrayQueueAttributes(t, "Compact queue", s, 0, lenExp, capExp)

	// test if all items in the queue are still in place

	for i := 0; i < s.Size(); i++ {
		if s.values[i] != i+1 {
			t.Fatalf("Compacted queue: item %v not found", i+1)
		}
	}

}

// Tests the queue add function.
func TestArrayQueueAdd(t *testing.T) {

	s := NewArrayQueue()

	// test empty queue attributes

	testArrayQueueAttributes(t, "Empty queue", s, 0, 0, initQueueCap)

	// test how add works

	lenExp := initQueueCap + 1
	capExp := initQueueCap + initQueueCap>>1

	for i := 1; i <= lenExp; i++ {
		s.Add(i)
	}

	// test full queue attributes

	testArrayQueueAttributes(t, "Full queue", s, 0, lenExp, capExp)

}

// Tests the queue peek function.
func TestArrayQueuePeek(t *testing.T) {

	s := NewArrayQueue()

	// test how peek works with empty queue

	i, ok := s.Peek()
	testArrayQueueResult(t, "Peek item from empty queue", i, nil, ok, false)
	testArrayQueueAttributes(t, "Peek item from empty queue", s, 0, 0, initQueueCap)

	// init queue

	lenExp := initQueueCap + 1
	capExp := initQueueCap + initQueueCap>>1

	for i := 1; i <= lenExp; i++ {
		s.Add(i)
	}

	// test how peek works with non-empty queue

	i, ok = s.Peek()
	testArrayQueueResult(t, "Peek item from non-empty queue", i, 1, ok, true)
	testArrayQueueAttributes(t, "Peek item from non-empty queue", s, 0, lenExp, capExp)

}

// Tests the queue poll function.
func TestArrayQueuePoll(t *testing.T) {

	s := NewArrayQueue()

	// init queue

	lenExp := initQueueCap
	capExp := initQueueCap

	for i := 1; i <= lenExp; i++ {
		s.Add(i)
	}

	// test how poll works with non-empty queue

	for i := 1; i <= lenExp; i++ {
		v, ok := s.Poll()
		testArrayQueueResult(t, "Poll item from non-empty queue", v, i, ok, true)
		testArrayQueueAttributes(t, "Poll item from empty queue", s, i, lenExp, capExp)
	}

	// test how poll works with empty queue

	i, ok := s.Poll()
	testArrayQueueResult(t, "Poll item from empty queue", i, nil, ok, false)
	testArrayQueueAttributes(t, "Poll item from empty queue", s, lenExp, lenExp, capExp)

}

// Tests the queue size function.
func TestArrayQueueSize(t *testing.T) {

	s := NewArrayQueue()

	// test empty queue size

	if s.Size() != 0 {
		t.Fatalf("Incorrect size of the queue, expected %v but was %v", 0, s.Size())
	}

	// init queue

	lenExp := initQueueCap

	for i := 1; i <= lenExp; i++ {
		s.Add(i)
	}

	// test non-empty queue size

	if s.Size() != lenExp {
		t.Fatalf("Incorrect size of the queue, expected %v but was %v", lenExp, s.Size())
	}

}

// Tests queue attributes if they equal expected values.
func testArrayQueueAttributes(t *testing.T, m string, s *ArrayQueue, headExp, tailExp, capExp int) {

	// test queue head and tail

	if s.head != headExp {
		t.Fatalf("%s: incorrect head of the queue, expected %v but was %v", m, headExp, s.head)
	}

	if s.tail != tailExp {
		t.Fatalf("%s: incorrect tail of the queue, expected %v but was %v", m, tailExp, s.tail)
	}

	// test queue underlying array length

	lenAct := len(s.values)
	if lenAct != capExp {
		t.Fatalf("%s: incorrect length of the queue underlying array, expected %v but was %v", m, capExp, lenAct)
	}

	// test queue underlying array capacity

	capAct := cap(s.values)
	if capAct != capExp {
		t.Fatalf("%s: incorrect capacity of the queue underlying array, expected %v but was %v", m, capExp, capAct)
	}

}

// Tests queue poll/peek function result values if they equal expected values.
func testArrayQueueResult(t *testing.T, m string, itemAct, itemExp interface{}, okAct, okExp bool) {

	if okAct != okExp {
		t.Fatalf("%s: OK flag value is incorrect, expected %v but was %v", m, okExp, okAct)
	}

	if itemAct != itemExp {
		t.Fatalf("%s: value of item taken from queue is incorrect, expected %v but was %v", m, itemExp, itemAct)
	}

}
