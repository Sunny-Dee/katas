package noderoutes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRoutes(t *testing.T) {
	var a, b, c, d, e Node

	a = Node{
		Data:     1,
		Children: []*Node{&b},
	}

	b = Node{
		Data:     2,
		Children: []*Node{&c, &d},
	}
	c = Node{
		Data:     3,
		Children: []*Node{&c, &d},
	}
	d = Node{
		Data:     4,
		Children: []*Node{},
	}
	e = Node{
		Data:     5,
		Children: []*Node{},
	}

	assert.True(t, FindRoute(&a, &c))
	assert.False(t, FindRoute(&a, &e))
}
