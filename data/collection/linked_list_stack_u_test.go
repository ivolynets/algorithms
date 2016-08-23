package collection

import "testing"

// Tests the stack push function.
func TestLinkedListStackPush(t *testing.T) {

	s := NewLinkedListStack()

	// test empty stack attributes

	testLinkedListStackAttributes(t, "Empty stack", s, 0, []int{})

	// test how push works

	valExp := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(valExp); i++ {
		s.Push(valExp[i])
	}

	// test full stack attributes

	testLinkedListStackAttributes(t, "Full stack", s, len(valExp), valExp)

}

// Tests the stack peek function.
func TestLinkedListStackPeek(t *testing.T) {

	s := NewLinkedListStack()

	// test how peek works with empty stack

	i, ok := s.Peek()
	testLinkedListStackResult(t, "Peek item from empty stack", i, nil, ok, false)
	testLinkedListStackAttributes(t, "Peek item from empty stack", s, 0, []int{})

	// init stack

	valExp := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(valExp); i++ {
		s.Push(valExp[i])
	}

	// test how peek works with non-empty stack

	i, ok = s.Peek()
	testLinkedListStackResult(t, "Peek item from non-empty stack", i, len(valExp), ok, true)
	testLinkedListStackAttributes(t, "Peek item from non-empty stack", s, len(valExp), valExp)

}

// Tests the stack pop function.
func TestLinkedListStackPop(t *testing.T) {

	s := NewLinkedListStack()

	// init stack

	valExp := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(valExp); i++ {
		s.Push(valExp[i])
	}

	// test how pop works with non-empty stack

	for len(valExp) > 0 {

		v, ok := s.Pop()
		testLinkedListStackResult(t, "Pop item from non-empty stack", v, valExp[len(valExp)-1], ok, true)

		valExp = valExp[:len(valExp)-1]
		testLinkedListStackAttributes(t, "Pop item from empty stack", s, len(valExp), valExp)
	}

	// test how pop works with empty stack

	i, ok := s.Pop()
	testLinkedListStackResult(t, "Pop item from empty stack", i, nil, ok, false)
	testLinkedListStackAttributes(t, "Pop item from empty stack", s, 0, []int{})

}

// Tests the stack size function.
func TestLinkedListStackSize(t *testing.T) {

	s := NewLinkedListStack()

	// test empty stack size

	if s.Size() != 0 {
		t.Fatalf("Incorrect size of the stack, expected %v but was %v", 0, s.Size())
	}

	// init stack

	valExp := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(valExp); i++ {
		s.Push(valExp[i])
	}

	// test non-empty stack size

	if s.Size() != len(valExp) {
		t.Fatalf("Incorrect size of the stack, expected %v but was %v", len(valExp), s.Size())
	}

}

// Tests stack attributes if they equal expected values.
func testLinkedListStackAttributes(t *testing.T, m string, s *LinkedListStack, sizeExp int, valExp []int) {

	// test stack size

	if s.size != sizeExp {
		t.Fatalf("%s: incorrect size of the stack, expected %v but was %v", m, sizeExp, s.size)
	}

	// test stack linked list values

	v := s.head
	for i := len(valExp); i > 0; i-- {

		if v == nil {
			t.Fatalf("%s: incorrect reference to the stack item, expected to be not %v but %v found", m, nil, nil)
		}

		if v.value != valExp[i-1] {
			t.Fatalf("%s: incorrect stack item, expected %v but was %v", m, valExp[i-1], v.value)
		}

		v = v.next

	}

}

// Tests stack pop/peek function result values if they equal expected values.
func testLinkedListStackResult(t *testing.T, m string, itemAct, itemExp interface{}, okAct, okExp bool) {

	if okAct != okExp {
		t.Fatalf("%s: OK flag value is incorrect, expected %v but was %v", m, okExp, okAct)
	}

	if itemAct != itemExp {
		t.Fatalf("%s: value of item taken from stack is incorrect, expected %v but was %v", m, itemExp, itemAct)
	}

}
