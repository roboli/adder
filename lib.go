// Package adds two numbers
package adder

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Adds two numbers
//
// Expects two int values, returns sum with an int value
// More: https://www.mathisfun.com/numbers/addition.html
func Add[T Number](a T, b T) T {
	return a + b
}
