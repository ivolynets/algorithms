package collection

// Stack general interface
type Stack interface {
	Push(i interface{})
	Peek() (interface{}, bool)
	Pop() (interface{}, bool)
	Size() int
}
