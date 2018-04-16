package sortstack

import (
	"testing"

	"github.com/golang-collections/collections/stack"
	"github.com/stretchr/testify/assert"
)

func TestSortStack(t *testing.T) {
	input := stack.New()
	input.Push(5)
	input.Push(2)
	input.Push(9)
	input.Push(3)

	expected := stack.New()
	expected.Push(9)
	expected.Push(5)
	expected.Push(3)
	expected.Push(2)

	actual := SortStack(input)

	for expected.Len() > 0 {

		assert.Equal(t, expected.Pop(), actual.Pop())
	}

}
