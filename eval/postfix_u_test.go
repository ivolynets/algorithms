package eval

import (
	"fmt"
	"testing"
)

// Tets the function for parsing mathematical expression into the postfix
// notation.
func TestPostfix(t *testing.T) {

	in := "3 + 4 * 2 / (1 - 5) ^ 2 ^ 3"

	outExp := "3 4 2 * 1 5 - 2 3 ^ ^ / +"
	outAct, err := Postfix(in)

	if err != nil {
		t.Fatalf("Function returned with error: %v", err)
	}

	if outAct != outExp {
		t.Fatalf("Incorrect output, expected %v but was %v", outExp, outAct)
	}

}

// Tets error case of the function for parsing mathematical expression into the
// postfix notation.
func TestPostfixError(t *testing.T) {

	in := "3 + mod(14.5, 7)"
	outAct, err := Postfix(in)

	if err == nil {
		t.Fatalf("Error expected but result returned: %v", outAct)
	}

	errExp := fmt.Sprintf("Invalid function or literal %q", "mod")
	if err.Error() != errExp {
		t.Fatalf("Invalid error message, expected %q but was %q", errExp, err.Error())
	}

}
