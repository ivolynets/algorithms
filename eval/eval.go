package eval

import (
	"fmt"
	"math"
	"strconv"

	"github.com/ivolynets/algorithms/data/collection"
)

// Evaluates an input mathematical expression string in postfix notation and
// returns result of that expression. It uses Reverse Polish Notation (RPN)
// algorithm by Charles L. Hamblin.
func EvalPostfix(in string) (float64, error) {

	// handles special cases

	if in == "" {
		return math.NaN(), fmt.Errorf("Empty expression")
	}

	s := collection.NewArrayStack()

	var f float64
	var err error
	var o func(collection.Stack) error

	c, b := "", ""

	for _, t := range in + " " {

		c = string(t)

		if c == " " {

			switch {

			case literal(b): // if the token is a literal, then push it onto the stack

				f, err = strconv.ParseFloat(b, 64)
				if err == nil {
					s.Push(f)
				} else {
					return math.NaN(), err
				}

			case operator(b): // if the token is an operator, then evaluate it

				if o, err = resolveOperator(b); err != nil {
					return math.NaN(), err
				} else {
					if err = o(s); err != nil {
						return math.NaN(), err
					}
				}

			case function(b): // if the token is a function, then evaluate it

				if o, err = resolveFunction(b); err != nil {
					return math.NaN(), err
				} else {
					if err = o(s); err != nil {
						return math.NaN(), err
					}
				}

			default:

				return math.NaN(), fmt.Errorf("Invalid symbol %q", b)

			}

			b = ""

		} else {
			b += c
		}

	}

	// in the end stack must contain exactly one item, the result of expression

	r, ok := s.Pop()

	if !ok {
		return math.NaN(), fmt.Errorf("Invalid postfix expression %q", in)
	}

	return r.(float64), nil

}

// Resolves operator by its symbol and returns corresponding function.
func resolveOperator(s string) (func(collection.Stack) error, error) {

	o, ok := operators[s]

	if ok {
		return o, nil
	}

	return nil, fmt.Errorf("Unknown operator %q", s)

}

// Resolves function by its name and returns corresponding function.
func resolveFunction(s string) (func(collection.Stack) error, error) {

	f, ok := functions[s]

	if ok {
		return f, nil
	}

	return nil, fmt.Errorf("Unknown operator %q", s)

}
