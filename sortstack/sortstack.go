package sortstack

import (
	"github.com/golang-collections/collections/stack"
)

// SortStack sorts a stack with the use of a helper stack.
func SortStack(s *stack.Stack) *stack.Stack {
	helper := stack.New()

	var current int
	for s.Len() > 0 {
		current = s.Pop().(int)
		for helper.Len() != 0 && helper.Peek().(int) < current {
			s.Push(helper.Pop())
		}
		helper.Push(current)
	}
	return helper
}
