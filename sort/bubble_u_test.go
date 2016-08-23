package sort

import (
	"reflect"
	"testing"

	"github.com/ivolynets/algorithms/data"
)

// Tests bubble sort function.
func TestBubbleSort(t *testing.T) {

	resAct := []data.Object{
		data.NewInt(3),
		data.NewInt(9),
		data.NewInt(4),
		data.NewInt(2),
		data.NewInt(1),
		data.NewInt(5),
		data.NewInt(7),
		data.NewInt(8),
		data.NewInt(6),
	}

	resExp := []data.Object{
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

	BubbleSort(resAct)

	if !reflect.DeepEqual(resAct, resExp) {
		t.Fatalf("Incorrect result, expected %v but was %v", resExp, resAct)
	}

}
