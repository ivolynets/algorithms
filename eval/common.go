package eval

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/ivolynets/algorithms/data/collection"
)

var operators = map[string]func(s collection.Stack) error{
	"!": opFact,
	"^": opPow,
	"*": opMul,
	"/": opDiv,
	"%": opMod,
	"+": opPlus,
	"-": opMinus,
}

var functions = map[string]func(s collection.Stack) error{
	"abs":  funcAbs,
	"acos": funcAcos,
	"acsc": funcAcsc,
	"actg": funcActg,
	"asec": funcAsec,
	"asin": funcAsin,
	"atan": funcAtan,
	"cos":  funcCos,
	"csc":  funcCsc,
	"ctg":  funcCtg,
	"lg":   funcLg,
	"ln":   funcLn,
	"log":  funcLog,
	"max":  funcMax,
	"min":  funcMin,
	"sec":  funcSec,
	"sin":  funcSin,
	"sqrt": funcSqrt,
	"tan":  funcTan,
}

// Determines if given string is a literal.
func literal(s string) bool {
	m, e := regexp.MatchString("[0-9]+(.[0-9]+)?", s)
	return e == nil && m
}

// Determines if given string is a variable.
func variable(s string) bool {
	m, e := regexp.MatchString("[a-zA-Z][0-9a-zA-Z_]*", s)
	return e == nil && m
}

// Determines if given string is an operator.
func operator(s string) bool {
	_, ok := operators[s]
	return ok
}

// Determines if given string is a function.
func function(s string) bool {
	_, ok := functions[s]
	return ok
}

// Returns the precedence of a given operator.
func prec(s string) int {

	switch {
	case strings.Contains("!", s):
		return 5
	case strings.Contains("^", s):
		return 4
	case strings.Contains("*/%", s):
		return 3
	case strings.Contains("+-", s):
		return 2
	}

	return 0

}

// Returns the associativity of a given operator. If operator has left
// associativity then function returns -1, if operator has right associativity
// then function returns 1, otherwise function returns 0.
func assoc(s string) int {

	switch {
	case strings.Contains("!^", s):
		return 1
	case strings.Contains("*/%+-", s):
		return -1
	}

	return 0

}

// Determines if given character is a left parenthesis.
func parLeft(s string) bool {
	return s == "("
}

// Determines if given character is a right parenthesis.
func parRight(s string) bool {
	return s == ")"
}

// Determines if given character is a function argument separator.
func separator(s string) bool {
	return s == ","
}

/*-- Operators ---------------------------------------------------------------*/

// Executes factorial operation.
func opFact(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Fact: missing argument")
	}

	// execute opertaion over arguments and store result

	s.Push(math.Gamma(x.(float64) + 1))
	return nil

}

// Executes power operation.
func opPow(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Pow: not enough arguments")
	}

	// execute opertaion over arguments and store result

	s.Push(math.Pow(x.(float64), y.(float64)))
	return nil

}

// Executes multiplication operation.
func opMul(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Mul: not enough arguments")
	}

	// execute opertaion over arguments and store result

	s.Push(x.(float64) * y.(float64))
	return nil

}

// Executes division operation.
func opDiv(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Div: not enough arguments")
	}

	// execute opertaion over arguments and store result

	s.Push(x.(float64) / y.(float64))
	return nil

}

// Executes modulus operation.
func opMod(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Mod: not enough arguments")
	}

	// execute opertaion over arguments and store result

	s.Push(math.Mod(x.(float64), y.(float64)))
	return nil

}

// Executes plus operation.
func opPlus(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Plus: not enough arguments")
	}

	// execute opertaion over arguments and store result

	s.Push(x.(float64) + y.(float64))
	return nil

}

// Executes minus operation.
func opMinus(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Minus: not enough arguments")
	}

	// execute opertaion over arguments and store result

	s.Push(x.(float64) - y.(float64))
	return nil

}

/*-- Functions ---------------------------------------------------------------*/

// Executes Abs function.
func funcAbs(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Abs: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Abs(x.(float64)))
	return nil

}

// Executes Acos function.
func funcAcos(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Acos: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Acos(x.(float64)))
	return nil

}

// Executes Acsc function.
func funcAcsc(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Acsc: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Asin(1 / x.(float64)))
	return nil

}

// Executes Actg function.
func funcActg(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Actg: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Atan(1 / x.(float64)))
	return nil

}

// Executes Asec function.
func funcAsec(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Asec: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Acos(1 / x.(float64)))
	return nil

}

// Executes Asin function.
func funcAsin(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Asin: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Asin(x.(float64)))
	return nil

}

// Executes Atan function.
func funcAtan(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Atan: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Atan(x.(float64)))
	return nil

}

// Executes Cos function.
func funcCos(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Cos: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Cos(x.(float64)))
	return nil

}

// Executes Csc function.
func funcCsc(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Csc: missing argument")
	}

	// execute function over arguments and store result

	s.Push(1 / math.Sin(x.(float64)))
	return nil

}

// Executes Ctg function.
func funcCtg(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Ctg: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Cos(x.(float64)) / math.Sin(x.(float64)))
	return nil

}

// Executes Lg function.
func funcLg(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Lg: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Log10(x.(float64)))
	return nil

}

// Executes Ln function.
func funcLn(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Ln: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Log(x.(float64)))
	return nil

}

// Executes Log function.
func funcLog(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Log: not enough arguments")
	}

	// execute function over arguments and store result

	s.Push(math.Log(y.(float64)) / math.Log(x.(float64)))
	return nil

}

// Executes Max function.
func funcMax(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Max: not enough arguments")
	}

	// execute function over arguments and store result

	s.Push(math.Max(x.(float64), y.(float64)))
	return nil

}

// Executes Min function.
func funcMin(s collection.Stack) error {

	// get and validate arguments

	y, ok := s.Pop()
	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Min: not enough arguments")
	}

	// execute function over arguments and store result

	s.Push(math.Min(x.(float64), y.(float64)))
	return nil

}

// Executes Sec function.
func funcSec(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Sec: missing argument")
	}

	// execute function over arguments and store result

	s.Push(1 / math.Cos(x.(float64)))
	return nil

}

// Executes Sin function.
func funcSin(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Sin: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Sin(x.(float64)))
	return nil

}

// Executes Sqrt function.
func funcSqrt(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Sqrt: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Sqrt(x.(float64)))
	return nil

}

// Executes Tan function.
func funcTan(s collection.Stack) error {

	// get and validate arguments

	x, ok := s.Pop()

	if !ok {
		return fmt.Errorf("Tan: missing argument")
	}

	// execute function over arguments and store result

	s.Push(math.Tan(x.(float64)))
	return nil

}
