package search

import (
	"testing"

	"github.com/ivolynets/algorithms/data"
)

// Tests the binary search function.
func TestBinarySearch(t *testing.T) {

	a := []data.Object{
		data.NewInt(1),
		data.NewInt(2),
		data.NewInt(3),
		data.NewInt(4),
		data.NewInt(5),
		data.NewInt(6),
		data.NewInt(7),
		data.NewInt(8),
		data.NewInt(9),
	}

	resExp := 6
	resAct := BinarySearch(a, data.NewInt(7))

	if resAct != resAct {
		t.Fatalf("Incorrect result, expected %v but was %v", resExp, resAct)
	}

}

// Tests the binary search of non-existing array element.
func TestBinarySearchFailure(t *testing.T) {

	a := []data.Object{
		data.NewInt(1),
		data.NewInt(2),
		data.NewInt(3),
		data.NewInt(4),
		data.NewInt(5),
		data.NewInt(6),
		data.NewInt(7),
		data.NewInt(8),
		data.NewInt(9),
	}

	resExp := -1
	resAct := BinarySearch(a, data.NewInt(10))

	if resAct != resAct {
		t.Fatalf("Incorrect result, expected %v but was %v", resExp, resAct)
	}

}

// Tests the binary search over the empty array.
func TestBinarySearchEmpty(t *testing.T) {

	a := []data.Object{}

	resExp := -1
	resAct := BinarySearch(a, data.NewInt(5))

	if resAct != resAct {
		t.Fatalf("Incorrect result, expected %v but was %v", resExp, resAct)
	}

}
