package data

// General object interface type
type Object interface {
	Compare(o Object) int
	Less(o Object) bool
	Greater(o Object) bool
	Equals(o Object) bool
	String() string
}
