// lib package used to perform arithmetic opeartions
//
// You can AddNums and SubNums
package workspace_lib

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// The function adds numbers
func AddNums(a, b int) int {
	return a + b
}

func Add[T Number](a, b T) T {
	return a + b
}

// This funciton subs numbers
func SubNums(a, b int) int {
	return a - b
}
