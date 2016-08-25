package eval

import (
	"fmt"
	"strings"

	"github.com/ivolynets/algorithms/data/collection"
)

// Function parses mathematical expression specified in infix notation into the
// one in postfix notation. It uses shunting-yard algorithm by Edsger Dijkstra.
func Postfix(in string) (out string, err error) {

	// cleanup input
	in = strings.Replace(in, " ", "", -1)

	out = ""
	err = nil

	s := collection.NewArrayStack()

	var i interface{} // container for stack items
	var c, o string   // container for characters
	var ok bool       // container for the success flag

	b := "" // buffer for literals and functions

	for _, t := range in {

		o = string(t)

		// if the token is an interceptor
		if b != "" && (operator(o) || parLeft(o) || parRight(o) || separator(o)) {

			switch {
			case literal(b):
				out += " " + b // if token is leteral, then add it to the output
			case function(b):
				s.Push(b) // if token is a function, then push it to stack
			default:
				err = fmt.Errorf("Invalid function or literal %q", b)
				return
			}

			// todo: handle variables

			b = ""

		}

		switch {

		case operator(o): // if token is an operator

			// while there is an operator at the top of the stack
			for i, ok = s.Peek(); ok && operator(i.(string)); i, ok = s.Peek() {

				c = i.(string)

				// o is left-associative and its predecence is less than or equal to that of c or
				// o is right-associative and its predecence is less than that of c
				if (assoc(o) == -1 && prec(o) <= prec(c)) || (assoc(o) == 1 && prec(o) < prec(c)) {

					// pop c off the operator stack onto the output
					out += " " + c
					s.Pop()

				} else {
					break
				}

			}

			// at the end of iteration push o onto the operator stack
			s.Push(o)

		case parLeft(o): // if the token is left parenthesis, then push it onto the stack

			s.Push(o)

		case parRight(o): // if the token is right parenthesis

			// until the token at the top of the stack is left parenthesis, pop operators off the stack onto the output queue
			// we also pop the left parenthesis from the stack, but not onto the output
			for i, ok = s.Pop(); ok && !parLeft(i.(string)); i, ok = s.Pop() {
				out += " " + i.(string)
			}

			// if the stack runs out without finding a left parenthesis, then there are mismatched parentheses
			if !ok {
				err = fmt.Errorf("Mismatched parentheses found")
				return
			}

			// if the token at the top of the stack is a function token
			i, ok = s.Peek()
			if ok && function(i.(string)) {

				// pop it onto the output queue
				out += " " + i.(string)
				s.Pop()

			}

		case separator(o): // if the token is a function argument separator

			// until the token at the top of the stack is left parenthesis, pop operators off the stack onto the output queue
			for i, ok = s.Peek(); ok && !parLeft(i.(string)); i, ok = s.Peek() {
				out += " " + i.(string)
				s.Pop()
			}

			// if no left parentheses are encountered, either the speparator was misplaced or parentheses were mismatched
			if !ok {
				err = fmt.Errorf("Separator is misplaced or mismatched parentheses found")
				return
			}

		default: // in all other cases add token to the buffer
			b += o

		}

	}

	// flush buffer
	if b != "" {
		if literal(b) {
			out += " " + b
		} else {
			err = fmt.Errorf("Invalid literal %q", b)
			return
		}
	}

	// while there are still operator tokens in the stack
	for i, ok = s.Pop(); ok; i, ok = s.Pop() {

		c = i.(string)

		// if the operator token at the top of the stack is a parenthesis, then there are mismatched parentheses
		if parLeft(c) || parRight(c) {
			err = fmt.Errorf("Mismatched parentheses found")
			return
		} else if operator(c) || function(c) {
			// pop the operator or function onto the output
			out += " " + c
		} else {
			err = fmt.Errorf("Invalid operator or function %q", c)
		}

	}

	return strings.TrimSpace(out), err

}
