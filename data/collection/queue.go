package collection

// Queue general interface
type Queue interface {
	Add(i interface{})
	Peek() (interface{}, bool)
	Poll() (interface{}, bool)
	Size() int
}
