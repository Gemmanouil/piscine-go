package piscine

type TreeNode00 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeLevelCount returns the number of levels (height) of the binary tree.
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)

	// Height is 1 (for current node) + max of left and right subtree heights
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
