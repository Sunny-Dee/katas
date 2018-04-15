package minimaltree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	description string
	input       []int
	expected    int
}{
	{
		description: "basic case",
		input:       []int{1, 2, 3, 4},
		expected:    3,
	},
	{
		description: "7 values case",
		input:       []int{1, 2, 3, 4, 8, 10, 11},
		expected:    3,
	},
}

func TestMinimalTree(t *testing.T) {
	for _, tc := range testCases {
		root := MinimalHeight(tc.input)
		actual := findHeight(root)
		assert.Equal(t, tc.expected, actual)
	}
}

func findHeight(n *TreeNode) int {
	height := 0
	q := []*TreeNode{n}
	fmt.Println(q)

	for len(q) != 0 {
		curr := q[0]
		if curr.left != nil {
			q = append(q, curr.left)
		}
		if curr.right != nil {
			q = append(q, curr.right)
		}

		q = q[1:]

		if curr.height > height {
			height = curr.height
		}
	}

	return height
}
