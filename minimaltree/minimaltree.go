package minimaltree

type TreeNode struct {
	data   int
	left   *TreeNode
	right  *TreeNode
	height int
}

// MinimalHeight returns the root node of a resulting balanced tree (shortest).
func MinimalHeight(arr []int) *TreeNode {
	return makeTree(arr, 0, len(arr)-1, 0)
}

func makeTree(arr []int, start, end, height int) *TreeNode {
	if end < start {
		return nil
	}
	var curr TreeNode

	mid := (end + start) / 2
	curr = TreeNode{
		data:   arr[mid],
		height: height + 1,
	}
	curr.left = makeTree(arr, start, mid-1, curr.height)
	curr.right = makeTree(arr, mid+1, end, curr.height)

	return &curr

}
