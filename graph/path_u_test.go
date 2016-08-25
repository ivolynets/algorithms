package graph

import (
	"reflect"
	"testing"
)

// Tests the function for finding the shortest path in the graph.
func TestPath(t *testing.T) {

	g := [][]float64{
		{0.0, 10.0, 8.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		{10.0, 0.0, 0.0, 6.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		{8.0, 0.0, 0.0, 0.0, 2.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		{0.0, 6.0, 0.0, 0.0, 2.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 2.0, 2.0, 0.0, 3.0, 1.0, 5.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0, 3.0, 0.0, 0.0, 0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0, 4.0},
		{0.0, 0.0, 0.0, 0.0, 5.0, 0.0, 0.0, 0.0, 7.0, 0.0},
		{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1.0, 7.0, 0.0, 0.0},
		{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 4.0, 0.0, 0.0, 0.0},
	}

	pathExp := []int{0, 2, 4, 6, 8}
	distExp := 12.0
	pathAct, distAct := Path(g, 0, 8)
	testPathResult(t, pathAct, pathExp, distAct, distExp)

	pathExp = []int{1, 3, 4, 2}
	distExp = 10.0
	pathAct, distAct = Path(g, 1, 2)
	testPathResult(t, pathAct, pathExp, distAct, distExp)

	pathExp = []int{7, 4, 6, 9}
	distExp = 10.0
	pathAct, distAct = Path(g, 7, 9)
	testPathResult(t, pathAct, pathExp, distAct, distExp)

}

// Tests path result by comparing expected values with actual ones.
func testPathResult(t *testing.T, pathAct, pathExp []int, distAct, distExp float64) {

	if !reflect.DeepEqual(pathAct, pathExp) {
		t.Fatalf("Incorrect path found, expected %v but was %v", pathExp, pathAct)
	}

	if distAct != distExp {
		t.Fatalf("Incorrect distance calculated, expected %v but was %v", distExp, distAct)
	}

}
