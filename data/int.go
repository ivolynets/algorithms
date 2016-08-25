package data

import "fmt"

// Integer wrapper type.
type Int struct {
	value int
}

// Creates a new integer and retunrs a pointer to it.
func NewInt(value int) *Int {
	return &Int{value}
}

// Returns integer value.
func (i *Int) Value() int {
	return i.value
}

// Compares two integeres, returns -1 if receiver is less than argument, 1 if
// receiver is greater than argument and 0 if they are equal.
func (i *Int) Compare(o Object) int {

	switch {
	case i.value < o.(*Int).value:
		return -1
	case i.value > o.(*Int).value:
		return 1
	default:
		return 0
	}

}

// Returns true if receiver is less than argument of false otherwise.
func (i *Int) Less(o Object) bool {
	return i.Compare(o) < 0
}

// Returns true if receiver is greater than argument of false otherwise.
func (i *Int) Greater(o Object) bool {
	return i.Compare(o) > 0
}

// Returns true if two integers are equal or false otherwise.
func (i *Int) Equals(o Object) bool {
	return i.Compare(o) == 0
}

// Returns string representation of integer.
func (i *Int) String() string {
	return fmt.Sprint(i.value)
}
