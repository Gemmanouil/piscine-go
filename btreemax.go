package piscine

type TreeNode02 struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeMax returns the node with the maximum value in the tree.
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
