package collection

import "testing"

// Test the stack internal extension function.
func TestArrayStackExtend(t *testing.T) {

	// init stack

	lenExp := 4
	capExp := 4

	s := &ArrayStack{make([]interface{}, capExp), 0}

	for i := 0; i < lenExp; i++ {
		s.values[i] = i + 1
		s.size++
	}

	// test if initial capacity of underlying array was not changed

	testArrayStackAttributes(t, "Initial stack", s, lenExp, capExp)

	// test how extension works

	capExp = 6
	s.extendStack(lenExp + 1)
	testArrayStackAttributes(t, "Extended stack", s, lenExp, capExp)

	// test if all items in the stack are still in place

	for i := 0; i < s.size; i++ {
		if s.values[i] != i+1 {
			t.Fatalf("Extended stack: item %v not found", i+1)
		}
	}

}

// Test the stack internal compaction function.
func TestArrayStackCompact(t *testing.T) {

	// init stack

	lenExp := initStackCap + 1
	capExp := initStackCap + initStackCap>>1
	capExp = capExp + capExp>>1

	s := &ArrayStack{make([]interface{}, capExp), 0}

	for i := 0; i < lenExp; i++ {
		s.values[i] = i + 1
		s.size++
	}

	// test if initial capacity of underlying array was not changed

	testArrayStackAttributes(t, "Initial stack", s, lenExp, capExp)

	// test how copaction works

	capExp = initStackCap + initStackCap>>1
	s.compactStack()
	testArrayStackAttributes(t, "Compact stack", s, lenExp, capExp)

	// test if all items in the stack are still in place

	for i := 0; i < s.size; i++ {
		if s.values[i] != i+1 {
			t.Fatalf("Compacted stack: item %v not found", i+1)
		}
	}

}

// Tests the stack push function.
func TestArrayStackPush(t *testing.T) {

	s := NewArrayStack()

	// test empty stack attributes

	testArrayStackAttributes(t, "Empty stack", s, 0, initStackCap)

	// test how push works

	lenExp := initStackCap + 1
	capExp := initStackCap + initStackCap>>1

	for i := 1; i <= lenExp; i++ {
		s.Push(i)
	}

	// test full stack attributes

	testArrayStackAttributes(t, "Full stack", s, lenExp, capExp)

}

// Tests the stack peek function.
func TestArrayStackPeek(t *testing.T) {

	s := NewArrayStack()

	// test how peek works with empty stack

	i, ok := s.Peek()
	testArrayStackResult(t, "Peek item from empty stack", i, nil, ok, false)
	testArrayStackAttributes(t, "Peek item from empty stack", s, 0, initStackCap)

	// init stack

	lenExp := initStackCap + 1
	capExp := initStackCap + initStackCap>>1

	for i := 1; i <= lenExp; i++ {
		s.Push(i)
	}

	// test how peek works with non-empty stack

	i, ok = s.Peek()
	testArrayStackResult(t, "Peek item from non-empty stack", i, lenExp, ok, true)
	testArrayStackAttributes(t, "Peek item from non-empty stack", s, lenExp, capExp)

}

// Tests the stack pop function.
func TestArrayStackPop(t *testing.T) {

	s := NewArrayStack()

	// init stack

	lenExp := initStackCap
	capExp := initStackCap

	for i := 1; i <= lenExp; i++ {
		s.Push(i)
	}

	// test how pop works with non-empty stack

	for i := lenExp; i > 0; i-- {
		v, ok := s.Pop()
		lenExp--
		testArrayStackResult(t, "Pop item from non-empty stack", v, i, ok, true)
		testArrayStackAttributes(t, "Pop item from empty stack", s, lenExp, capExp)
	}

	// test how pop works with empty stack

	i, ok := s.Pop()
	testArrayStackResult(t, "Pop item from empty stack", i, nil, ok, false)
	testArrayStackAttributes(t, "Pop item from empty stack", s, 0, capExp)

}

// Tests the stack size function.
func TestArrayStackSize(t *testing.T) {

	s := NewArrayStack()

	// test empty stack size

	if s.Size() != 0 {
		t.Fatalf("Incorrect size of the stack, expected %v but was %v", 0, s.Size())
	}

	// init stack

	lenExp := initStackCap

	for i := 1; i <= lenExp; i++ {
		s.Push(i)
	}

	// test non-empty stack size

	if s.Size() != lenExp {
		t.Fatalf("Incorrect size of the stack, expected %v but was %v", lenExp, s.Size())
	}

}

// Tests stack attributes if they equal expected values.
func testArrayStackAttributes(t *testing.T, m string, s *ArrayStack, sizeExp, capExp int) {

	// test stack size

	if s.size != sizeExp {
		t.Fatalf("%s: incorrect size of the stack, expected %v but was %v", m, sizeExp, s.size)
	}

	// test stack underlying array length

	lenAct := len(s.values)
	if lenAct != capExp {
		t.Fatalf("%s: incorrect length of the stack underlying array, expected %v but was %v", m, capExp, lenAct)
	}

	// test stack underlying array capacity

	capAct := cap(s.values)
	if capAct != capExp {
		t.Fatalf("%s: incorrect capacity of the stack underlying array, expected %v but was %v", m, capExp, capAct)
	}

}

// Tests stack pop/peek function result values if they equal expected values.
func testArrayStackResult(t *testing.T, m string, itemAct, itemExp interface{}, okAct, okExp bool) {

	if okAct != okExp {
		t.Fatalf("%s: OK flag value is incorrect, expected %v but was %v", m, okExp, okAct)
	}

	if itemAct != itemExp {
		t.Fatalf("%s: value of item taken from stack is incorrect, expected %v but was %v", m, itemExp, itemAct)
	}

}
