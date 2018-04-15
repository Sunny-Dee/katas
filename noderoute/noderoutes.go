package noderoutes

import (
	"github.com/golang-collections/collections/stack"
)

// Represents a node in a directed graph
type Node struct {
	Data     int
	Children []*Node
	Visited  bool
}

// FindRoute finds if there is a route between two nodes
func FindRoute(a, b *Node) bool {
	stck := stack.New()
	stck.Push(a)
	var curr *Node

	for stck.Len() > 0 {
		curr = stck.Pop().(*Node)
		for _, child := range curr.Children {
			if !child.Visited {
				if child.Data == b.Data {
					return true
				}
				child.Visited = true
				stck.Push(child)
			}
		}
	}

	return false
}
