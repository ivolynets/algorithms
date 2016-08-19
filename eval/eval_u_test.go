package eval

import (
	"fmt"
	"testing"
)

// Tests the function for evaluation of mathematical expression in postfix
// notation.
func TestEvalPostfix(t *testing.T) {

	in := "3 4.5 sin ! 2 * 7 5 3 2 min % - 2 3 ^ ^ / +" // 3 + sin(4.5)! * 2 / (7 - (5 % min(3, 2))) ^ 2 ^ 3 = 3.0000523316959189505553006102527

	resExp := 3.0000523316959189505553006102527
	resAct, err := EvalPostfix(in)

	if err != nil {
		t.Fatalf("Function returned with error: %v", err)
	}

	if resAct != resExp {
		t.Fatalf("Incorrect result, expected %v but was %v", resExp, resAct)
	}

}

// Tests error case of the function for evaluation of mathematical expression
// in postfix notation.
func TestEvalPostfixError(t *testing.T) {

	in := "3 4.5 7 mod +" // 3 + mod(4.5, 7)
	resAct, err := EvalPostfix(in)

	if err == nil {
		t.Fatalf("Error expected but result returned: %v", resAct)
	}

	errExp := fmt.Sprintf("Invalid symbol %q", "mod")
	if err.Error() != errExp {
		t.Fatalf("Invalid error message, expected %q but was %q", errExp, err.Error())
	}

}
